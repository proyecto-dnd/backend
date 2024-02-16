package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/proyecto-dnd/backend/cmd/server/router"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/pkg/firebaseConnection"

	"firebase.google.com/go/v4/auth"
	// "firebase.google.com/go/v4/auth"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := ConnectDB()
	firebaseApp := firebaseConnection.InitializeFirebaseApp()

	engine := gin.New()
	engine.Use(gin.Recovery())

	router := router.NewRouter(engine, db, firebaseApp)
	router.MapRoutes()
	// createOneuser()

	if err := engine.Run("localhost:8080"); err != nil {
		panic(err)
	}
	defer db.Close()
}

func createOneuser() {
	firebaseClient := firebaseConnection.CreateFirebaseClient()

	ctx := &gin.Context{}

	userParams := (&auth.UserToCreate{}).
		UID("123qweasd").
		Email("dthmax2@gmail.com").
		Password("pass123").
		DisplayName("Kai").
		Disabled(false)
	u, err := firebaseClient.CreateUser(ctx, userParams)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	var userTemp domain.User
	userTemp.Username = u.DisplayName
	userTemp.Email = u.Email

	log.Printf("Successfully created user: %v\n", u)
	log.Printf("User Info is: %v\n", userTemp)
}

// func initializeFirebaseApp() *firebase.App {
// 	ctx := &gin.Context{}
// 	opt := option.WithCredentialsFile("../serviceAccountKey.json")

// 	conf := &firebase.Config{
// 		ProjectID:     "logoflegends-19383",
// 		StorageBucket: "logoflegends-19383.appspot.com",
// 	}

// 	app, err := firebase.NewApp(ctx, conf, opt)
// 	if err != nil {
// 		log.Fatalf("error initializing app: %v\n", err)
// 	}

// 	// firestoreClient, err := firebaseApp.Firestore(ctx)
// 	// if err != nil {
// 	// 	log.Fatalf("error creating client: %v\n", err)
// 	// }

// 	// ref := firestoreClient.Collection("testDB").NewDoc()
// 	// result, err := ref.Set(ctx, map[string]interface{}{
// 	// 	"title":       "la primera...",
// 	// 	"description": "Y la segunda...",
// 	// })
// 	// if err != nil {
// 	// 	log.Fatalf("error when creating testDB: %v\n", err)
// 	// }

// 	// log.Printf("Result is: [%v]", result)
// 	fmt.Println("conectado a firebase")
// 	return app
// }

func ConnectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")

	// Create the data source.
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open the connection.
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	// Check the connection.
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Se conectó a la base de datos bien")
	return db
}
