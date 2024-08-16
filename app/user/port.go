package user

import (
	"context"

	userProto "github.com/promptlabth/proto-lib/user"
)

type userAdaptor interface {
	GetDetailUser(ctx context.Context, firebaseId string) (*userProto.GetUserByIdRes, error)
	UpsertUser(ctx context.Context, req *userProto.UpsertUserReq) (*userProto.UpsertUserRes, error)
}

type stripeAdaptor interface {
}

type firebaseAdaptor interface {
}
