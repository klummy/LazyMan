package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// FirebaseApp - Instance of Firebase Client
func FirebaseApp(credentialsPath string) (*firebase.App, error) {
	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return nil, err
	}

	return app, nil
}

// VerifyIDToken - Validate a given Firebase client
func VerifyIDToken(ctx context.Context, idToken string, firebaseApp *firebase.App) (*auth.Token, error) {
	client, clientErr := firebaseApp.Auth(ctx)
	if clientErr != nil {
		log.Println("Error getting Firebase auth client: ", clientErr)
		return nil, clientErr
	}

	token, tokenErr := client.VerifyIDToken(idToken)
	if tokenErr != nil {
		log.Println("Error verifying Firebase token: ", tokenErr)
		return nil, tokenErr
	}

	return token, nil
}
