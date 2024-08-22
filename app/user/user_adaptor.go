package user

import (
	"context"
	"time"

	"github.com/promptlabth/ms-ai-marketplace/logger"
	userProto "github.com/promptlabth/proto-lib/user"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc/metadata"
)

func (a *UserAdaptor) GetDetailUser(ctx context.Context, firebaseId string) (*userProto.GetUserByIdRes, error) {
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
	)
	ctxReq := metadata.NewOutgoingContext(ctx, md)
	otelgrpc.Inject(ctxReq, &md)
	res, err := a.userServiceClient.GetDetailUser(ctxReq, &userProto.GetUserByIdReq{
		FirebaseId: firebaseId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *UserAdaptor) UpsertUser(ctx context.Context, req *userProto.UpsertUserReq) (*userProto.UpsertUserRes, error) {
	logger.Info(ctx, "Request to Upsert User")
	res, err := a.userServiceClient.UpsertUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
