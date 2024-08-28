// usecase.go

package user

import (
	"context"
	"strconv"

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

// UserUsecase struct that combines storage and domain to execute user-related business logic.
type UserUsecase struct {
	storage        userRepository
	userAdaptor    userAdaptor
	grpcUserServer grpcUserServer
}

// NewUsecase creates a new Usecase instance with the provided storage and domain logic implementations.
func NewUsecase(s userRepository, userAdaptor userAdaptor, userServer grpcUserServer) *UserUsecase {
	return &UserUsecase{
		storage:        s,
		userAdaptor:    userAdaptor,
		grpcUserServer: userServer,
	}
}

func (u *UserUsecase) GetUser(ctx context.Context, firebaseID string) (*UserEntity, error) {
	user, err := u.grpcUserServer.GetDetailUser(ctx, firebaseID)
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
