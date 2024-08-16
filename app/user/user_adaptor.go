package user

import (
	"context"

	userProto "github.com/promptlabth/proto-lib/user"
)

func (a *UserAdaptor) GetDetailUser(ctx context.Context, firebaseId string) (*userProto.GetUserByIdRes, error) {
	res, err := a.userServiceClient.GetDetailUser(ctx, &userProto.GetUserByIdReq{
		FirebaseId: firebaseId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *UserAdaptor) UpsertUser(ctx context.Context, req *userProto.UpsertUserReq) (*userProto.UpsertUserRes, error) {
	res, err := a.userServiceClient.UpsertUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
