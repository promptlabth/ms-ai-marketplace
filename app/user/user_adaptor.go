package user

import (
	"context"
	"time"

	"github.com/promptlabth/ms-ai-marketplace/logger"
	userProto "github.com/promptlabth/proto-lib/user"
	"google.golang.org/grpc/metadata"
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
	logger.Info(ctx, "Request to Upsert User")
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
	)
	ctxReq := metadata.NewOutgoingContext(ctx, md)
	res, err := a.userServiceClient.UpsertUser(ctxReq, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
