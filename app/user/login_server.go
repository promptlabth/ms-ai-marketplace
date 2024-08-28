package user

import (
	"context"
	"time"

	"github.com/promptlabth/ms-ai-marketplace/logger"
	userProto "github.com/promptlabth/proto-lib/user"
	"google.golang.org/grpc/metadata"
)

func (a *UserServer) UpsertUser(ctx context.Context, req *userProto.UpsertUserReq) (*userProto.UpsertUserRes, error) {
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
