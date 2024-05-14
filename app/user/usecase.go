// usecase.go

package user

import (
	"context"
	"time"
)

// Assuming User struct is defined in user.go within the 'user' package.

// storage outlines the methods required by the use case to interact with the data layer.
type storage interface {
	CreateUser(context.Context, UserEntity) (*string, error)
	// Add other storage methods as necessary
}

// domain outlines the methods required by the use case for domain logic and validations.
type domain interface {
	ValidateNewUser(ctx context.Context, user User) error
	// Add other domain methods as necessary
}

// Usecase struct that combines storage and domain to execute user-related business logic.
type Usecase struct {
	storage storage
	domain  domain
}

// NewUsecase creates a new Usecase instance with the provided storage and domain logic implementations.
func NewUsecase(s storage, d domain) *Usecase {
	return &Usecase{
		storage: s,
		domain:  d,
	}
}

// NewUser orchestrates the process of validating and creating a new user.
func (u *Usecase) NewUser(ctx context.Context, user User) error {
	// First, use the domain logic to validate the new user.
	if err := u.domain.ValidateNewUser(ctx, user); err != nil {
		return err
	}

	userEntity := UserEntity{
		FriebaseID:     user.FriebaseID,
		Name:           user.Name,
		Email:          user.Email,
		Platform:       user.Platform,
		PlanID:         user.PlanID,
		StripeID:       user.StripeID,
		LastActiveTime: time.Now(),
	}

	// If validation passes, proceed to create the user in the storage layer.
	_, err := u.storage.CreateUser(ctx, userEntity)
	return err
}
