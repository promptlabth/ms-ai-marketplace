// usecase.go

package user

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/promptlabth/ms-ai-marketplace/app"
)

// Assuming User struct is defined in user.go within the 'user' package.

// userRepository outlines the methods required by the use case to interact with the data layer.
type userRepository interface {
	CreateUser(context.Context, UserEntity) (*string, error)
	GetUserByID(context.Context, string) (*UserEntity, error)
	UpdateUser(context.Context, UserEntity) error
	DeleteUser(context.Context, string) error
}

// domain outlines the methods required by the use case for domain logic and validations.
type domain interface {
	ValidateNewUser(ctx context.Context, user User) error
	// Add other domain methods as necessary
}

// UserUsecase struct that combines storage and domain to execute user-related business logic.
type UserUsecase struct {
	storage     userRepository
	domain      domain
	userAdaptor userAdaptor
}

// NewUsecase creates a new Usecase instance with the provided storage and domain logic implementations.
func NewUsecase(s userRepository, d domain, userAdaptor userAdaptor) *UserUsecase {
	return &UserUsecase{
		storage:     s,
		domain:      d,
		userAdaptor: userAdaptor,
	}
}

// NewUser orchestrates the process of validating and creating a new user.
func (u *UserUsecase) NewUser(ctx context.Context, user User) error {

	// First, use the domain logic to validate the new user.
	if err := u.domain.ValidateNewUser(ctx, user); err != nil {
		log.Printf("Error validating new user: %v", err)
		return err
	}

	userEntity := UserEntity{
		FirebaseID:     user.FirebaseID,
		Name:           user.Name,
		Email:          user.Email,
		Platform:       user.Platform,
		PlanID:         user.PlanID,
		StripeID:       user.StripeID,
		LastActiveTime: time.Now(),
	}
	log.Printf("UserEntity : %+v\n", userEntity)

	// If validation passes, proceed to create the user in the storage layer.
	_, err := u.storage.CreateUser(ctx, userEntity)
	return err
}

func (u *UserUsecase) GetUser(ctx context.Context, firebaseID string) (*UserEntity, error) {
	user, err := u.userAdaptor.GetDetailUser(ctx, firebaseID)
	if err != nil {
		return nil, err
	}

	return &UserEntity{
		FirebaseID: user.FirebaseId,
		Name:       user.Name,
		Email:      app.PtrToType(user.Email),
		Platform:   app.PtrToType(user.Platform),
		StripeID:   app.PtrToType(user.StripeId),
		PlanID:     strconv.Itoa(int(user.Plan.Id)),
	}, nil
}
