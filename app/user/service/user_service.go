package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/promptlabth/ms-ai-marketplace/app/user/repository"
)

type contextKey string

const maxMessagesKey contextKey = "maxMessages"

var jwtSecret = []byte("your-secret-key")

type userService struct {
	userRepository repository.UserRepository
}

func callExternalAPI(url string, method string, payload interface{}, headers map[string]string) (*http.Response, error) {
	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Create an HTTP client and set a timeout
	client := &http.Client{Timeout: 10 * time.Second}

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	return resp, nil
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (s userService) NewUser(ctx context.Context, request NewUserRequest) (*UserResponse, error) {
	// Check if the user already exists
	url := os.Getenv("PROMPTLAB_MAIN")
	payload := map[string]interface{}{
		"platform":     "gmail",
		"access_token": "",
	}

	headers := map[string]string{
		"Authorization": "Bearer " + request.AccessToken,
	}
	Promplab_res, err := callExternalAPI(url, "POST", payload, headers)
	if err != nil {
		log.Printf("Error calling external API: %v", err)
		// Handle the error as needed
		return nil, err
	}
	if Promplab_res != nil {
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
		var promplabResponse PromplabResponse
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
			PlanID:         updatedUser.PlanID,
			ProfilePicture: updatedUser.ProfilePicture,
			AccessToken:    updatedUser.AccessToken,
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
	}

	return &response, nil
}


func (s userService) GetUser(firebaseID string) (UserResponse, error) {
	user, err := s.userRepository.GetUserByFirebaseID(firebaseID)
	if err != nil {
		log.Println(err) // Use log.Println instead of log.Fatal
		return UserResponse{}, err
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
