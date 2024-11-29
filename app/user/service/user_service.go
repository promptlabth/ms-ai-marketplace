package service

import (
	"log"
	"time"

	"github.com/promptlabth/ms-ai-marketplace/app/user/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return userService{userRepository: userRepository}
}

func (s userService) NewUser(request NewUserRequest) (*UserResponse, error) {
	// Check if the user already exists
	existingUser, err := s.userRepository.GetUserByFirebaseID(request.FirebaseID)
	if (err == nil && existingUser != nil) {
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
