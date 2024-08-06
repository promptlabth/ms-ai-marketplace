package user

import (
	"context"
)

func (a *UserAdaptor) GetDetailUser(ctx context.Context, firebaseId string) (*GetUserByIdRes, error) {
	res, err := a.userServiceClient.GetDetailUser(ctx, &GetUserByIdReq{
		FirebaseId: firebaseId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *UserAdaptor) UpsertUser(ctx context.Context, req *UpsertUserReq) (*UpsertUserRes, error) {
	res, err := a.userServiceClient.UpsertUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
