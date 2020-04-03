package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"firebase.google.com/go/storage"
	"log"
)

func NewFirestoreClient() *firestore.Client {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing client: %v\n", err)
	}

	return client
}

func NewFirebaseAuthClient() *auth.Client {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing client: %v\n", err)
	}

	return client
}

func NewCloudMessagingClient() *messaging.Client {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error initializing client: %v\n", err)
	}

	return client
}

func NewCloudStorageClient() *storage.Client {
	ctx := context.Background()

	config := &firebase.Config{
		StorageBucket: "covidlocal-staging.appspot.com",
	}

	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("error initializing client: %v\n", err)
	}

	return client
}
