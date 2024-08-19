package user

import (
	"context"

	auth "firebase.google.com/go/v4/auth"
	userProto "github.com/promptlabth/proto-lib/user"
)

type userAdaptor interface {
	GetDetailUser(ctx context.Context, firebaseId string) (*userProto.GetUserByIdRes, error)
	UpsertUser(ctx context.Context, req *userProto.UpsertUserReq) (*userProto.UpsertUserRes, error)

	ValidateToken(ctx context.Context, tokenId string) (*auth.Token, error)
}
