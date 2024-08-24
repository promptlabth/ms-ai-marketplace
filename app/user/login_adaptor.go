package user

import (
	"context"

	auth "firebase.google.com/go/v4/auth"
	"github.com/promptlabth/ms-ai-marketplace/logger"
)

func (a *UserAdaptor) ValidateToken(ctx context.Context, tokenId string) (*auth.Token, error) {
	logger.Info(ctx, "Request to auth for firebase")
	client, err := a.firebase.Auth(ctx)
	if err != nil {
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, tokenId)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	return token, err
}

func (a *UserAdaptor) FirebaseRetrieveUserData(ctx context.Context, firebaseId string) (*UserDetailDomain, error) {
	client, err := a.firebase.Auth(ctx)
	if err != nil {
		return nil, err
	}

	usr, err := client.GetUser(ctx, firebaseId)
	if err != nil {
		return nil, err
	}

	return &UserDetailDomain{
		FirebaseId:    firebaseId,
		Name:          usr.DisplayName,
		ProfilePicUrl: usr.PhotoURL,
		Email:         usr.Email,
	}, nil
}
