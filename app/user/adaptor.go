// adaptor.go

package user

import (
	firebase "firebase.google.com/go/v4"
	userProto "github.com/promptlabth/proto-lib/user"
)

type UserAdaptor struct {
	userServiceClient userProto.UserServiceClient
	firebase          *firebase.App
}

func NewUserAdaptor(userServiceClient userProto.UserServiceClient) *UserAdaptor {
	return &UserAdaptor{
		userServiceClient: userServiceClient,
	}
}
