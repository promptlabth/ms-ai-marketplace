package auth

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func Init(ctx context.Context) (*firebase.App, error) {
	opt := option.WithCredentialsFile("")
	config := &firebase.Config{ProjectID: ""}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}
