package user

import (
	"context"

	userProto "github.com/promptlabth/proto-lib/user"
)

type UserServer struct {
	userServiceClient userProto.UserServiceClient
}

func NewGrpcServer(userServiceClient userProto.UserServiceClient) *UserServer {
	return &UserServer{
		userServiceClient: userServiceClient,
	}
}

func (a *UserServer) GetDetailUser(ctx context.Context, firebaseId string) (*userProto.GetUserByIdRes, error) {
	res, err := a.userServiceClient.GetDetailUser(ctx, &userProto.GetUserByIdReq{
		FirebaseId: firebaseId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
