package auth

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func Init() (*firebase.App, error) {
	opt := option.WithCredentialsFile("firebase-credential.json")
	config := &firebase.Config{ProjectID: "prompt-lab-383408"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}
