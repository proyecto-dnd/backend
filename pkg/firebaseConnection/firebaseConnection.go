package firebaseConnection

import (
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var (
	ctx = &gin.Context{}
)

func InitializeFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile("../serviceAccountKey.json")

	conf := &firebase.Config{
		ProjectID:     "logoflegends-19383",
		StorageBucket: "logoflegends-19383.appspot.com",
	}

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// firestoreClient, err := firebaseApp.Firestore(ctx)
	// if err != nil {
	// 	log.Fatalf("error creating client: %v\n", err)
	// }

	// ref := firestoreClient.Collection("testDB").NewDoc()
	// result, err := ref.Set(ctx, map[string]interface{}{
	// 	"title":       "la primera...",
	// 	"description": "Y la segunda...",
	// })
	// if err != nil {
	// 	log.Fatalf("error when creating testDB: %v\n", err)
	// }

	// log.Printf("Result is: [%v]", result)
	fmt.Println("conectado a firebase")
	return app
}

func CreateFirebaseClient() *auth.Client {
	app := InitializeFirebaseApp()

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return client
}
