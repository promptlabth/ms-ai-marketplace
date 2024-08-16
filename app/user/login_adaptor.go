package user

import (
	"context"

	auth "firebase.google.com/go/v4/auth"
)

func (a *UserAdaptor) ValidateToken(ctx context.Context, tokenId string) (*auth.Token, error) {
	client, err := a.firebase.Auth(ctx)
	if err != nil {
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, tokenId)
	if err != nil {
		return nil, err
	}
	return token, err
}
