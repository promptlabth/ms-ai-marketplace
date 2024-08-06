package user

import "context"

type userAdaptor interface {
	GetDetailUser(ctx context.Context, firebaseId string) (*GetUserByIdRes, error)
	UpsertUser(ctx context.Context, req *UpsertUserReq) (*UpsertUserRes, error)
}
