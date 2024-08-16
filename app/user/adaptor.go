// adaptor.go

package user

import userProto "github.com/promptlabth/proto-lib/user"

type UserAdaptor struct {
	userServiceClient userProto.UserServiceClient
}

func NewUserAdaptor(userServiceClient userProto.UserServiceClient) *UserAdaptor {
	return &UserAdaptor{
		userServiceClient: userServiceClient,
	}
}
