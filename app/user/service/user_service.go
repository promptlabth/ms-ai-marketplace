package service

import (
	// "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/promptlabth/ms-ai-marketplace/app/user/repository"
	"github.com/promptlabth/ms-ai-marketplace/app/utils"
)

type contextKey string

const maxMessagesKey contextKey = "maxMessages"

var jwtSecret = []byte("your-secret-key")


type userService struct {
	userRepository repository.UserRepository
	AccessTokenPromptlab    string
}


func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) NewUser(ctx context.Context, request NewUserRequest) (*UserResponse, error) {
    s.AccessTokenPromptlab = request.AccessToken

    // Check if the user already exists
    path := "/v1/login"
    // url := os.Getenv("PROMPTLAB_MAIN")
	url := "https://prompt-lab-be-uu4qhhj35a-as.a.run.app"

    if url == "" {
        return nil, errors.New("PROMPTLAB_MAIN environment variable is not set")
    }

    payload := map[string]interface{}{
        "platform":     "gmail",
        "access_token": "",
    }

    headers := map[string]string{
        "Authorization": "Bearer " + request.AccessToken,
    }

    // Declare promplabResponse outside of the if block
    var promplabResponse PromplabResponse

    Promplab_res, err := utils.CallExternalAPI(url, "POST", payload, headers, path)
    if err != nil {
        log.Printf("Error calling external API: %v", err)
        return nil, err
    }

    // Handle the response from the external API
    if (Promplab_res != nil) {
        defer Promplab_res.Body.Close()
        if Promplab_res.StatusCode == http.StatusUnauthorized {
            log.Printf("Received 401 Unauthorized response code")
            return nil, errors.New("received 401 Unauthorized response code")
        }
        if Promplab_res.StatusCode != http.StatusOK {
            log.Printf("Received non-200 response code: %d", Promplab_res.StatusCode)
            return nil, fmt.Errorf("received non-200 response code: %d", Promplab_res.StatusCode)
        }
        body, err := io.ReadAll(Promplab_res.Body)
        if err != nil {
            log.Printf("Error reading response body: %v", err)
            return nil, err
        }
        err = json.Unmarshal(body, &promplabResponse)
        if err != nil {
            log.Printf("Error unmarshalling response body: %v", err)
            return nil, err
        }

        fmt.Printf("Max Messages: %d\n", promplabResponse.Plan.Product.MaxMessages)

        ctx = context.WithValue(ctx, maxMessagesKey, promplabResponse.Plan.Product.MaxMessages)
        printContextValue(ctx)
    }

    existingUser, err := s.userRepository.GetUserByFirebaseID(request.FirebaseID)
    if err == nil && existingUser != nil {
        // User exists, update DatetimeLastActive
        existingUser.DatetimeLastActive = time.Now().Format(time.RFC3339)
        existingUser.PlanID = promplabResponse.Plan.Product.PlanType
        existingUser.MaxMessages = promplabResponse.Plan.Product.MaxMessages
        updatedUser, err := s.userRepository.Update(*existingUser)
        if err != nil {
            log.Fatal(err)
            return nil, err
        }

        response := UserResponse{
            ID:             updatedUser.ID,
            FirbaseID:      updatedUser.FirebaseID,
            Name:           updatedUser.Name,
            Email:          updatedUser.Email,
            Platform:       updatedUser.Platform,
            PlanID:         promplabResponse.Plan.Product.PlanType, // Correct field name
            ProfilePicture: updatedUser.ProfilePicture,
            AccessToken:    updatedUser.AccessToken,
            MaxMessages:    updatedUser.MaxMessages,
        }

        return &response, nil
    }

    // User does not exist, create a new user
    user := repository.User{
        FirebaseID:         request.FirebaseID,
        Name:               request.Name,
        Email:              request.Email,
        Platform:           request.Platform,
        StripeID:           request.StripeID,
        PlanID:             request.PlanID,
        DatetimeLastActive: time.Now().Format(time.RFC3339),
        ProfilePicture:     request.ProfilePicture,
        AccessToken:        request.AccessToken,
        MaxMessages: 	  promplabResponse.Plan.Product.MaxMessages,
    }

    newUser, err := s.userRepository.Create(user)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    response := UserResponse{
        ID:             newUser.ID,
        FirbaseID:      newUser.FirebaseID,
        Name:           newUser.Name,
        Email:          newUser.Email,
        Platform:       newUser.Platform,
        PlanID:         newUser.PlanID,
        ProfilePicture: newUser.ProfilePicture,
        AccessToken:    newUser.AccessToken,
        MaxMessages:  	newUser.MaxMessages,

    }

    return &response, nil
}

func (s *userService) GetUser(firebaseID string) (UserResponse, error) {
    user, err := s.userRepository.GetUserByFirebaseID(firebaseID)
    if err != nil {
        log.Println(err) // Use log.Println instead of log.Fatal
        return UserResponse{}, err
    }
    // fmt.Println("Promplab Access Token: ", s.AccessTokenPromptlab)
    path := "/v1/user/remaining-message"
    // url := os.Getenv("PROMPTLAB_MAIN")
	url := "https://prompt-lab-be-uu4qhhj35a-as.a.run.app"

    headers := map[string]string{
        "Authorization": "Bearer " + s.AccessTokenPromptlab,
    }
    Promplab_res, err := utils.CallExternalAPI(url, "GET", nil, headers, path) // Pass nil as payload
    if err != nil {
        log.Printf("Error calling external API: %v", err)
        return UserResponse{}, err
    }
    fmt.Println("Promplab Response: ", Promplab_res)

    var usedMessages int
    if Promplab_res != nil {
        defer Promplab_res.Body.Close()
        if Promplab_res.StatusCode == http.StatusOK {
            body, err := io.ReadAll(Promplab_res.Body)
            if err != nil {
                log.Printf("Error reading response body: %v", err)
                return UserResponse{}, err
            }
            fmt.Println("Response Body: ", string(body)) // Print the response body for debugging

            // Check if the response body is a number
            var responseNumber float64
            err = json.Unmarshal(body, &responseNumber)
            if err == nil {
                usedMessages = int(responseNumber)
            } else {
                // If not a number, try to unmarshal into a map
                var responseMap map[string]interface{}
                err = json.Unmarshal(body, &responseMap)
                if err != nil {
                    log.Printf("Error unmarshalling response body: %v", err)
                    return UserResponse{}, err
                }
                if usedMessagesValue, ok := responseMap["used_messages"].(float64); ok {
                    usedMessages = int(usedMessagesValue)
                }
            }
        }
    }

    response := UserResponse{
        ID:             user.ID,
        FirbaseID:      user.FirebaseID,
        Name:           user.Name,
        Email:          user.Email,
        Platform:       user.Platform,
        PlanID:         user.PlanID,
        ProfilePicture: user.ProfilePicture,
        AccessToken:    user.AccessToken,
        Role:           user.Role,
        MaxMessages:    user.MaxMessages,
        UsedMessages:   usedMessages,
    }

    return response, nil
}

func GenerateJWT(firebaseID string) (string, error) {
	claims := jwt.MapClaims{
		"firebase_id": firebaseID,
		"exp":         time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func printContextValue(ctx context.Context) {
    maxMessages, ok := ctx.Value(maxMessagesKey).(int)
    if !ok {
        fmt.Println("MaxMessages not found in context")
        return
    }
    fmt.Printf("Context: MaxMessages: %d\n", maxMessages)
}
