package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main () {
	// Recover from panic.
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// Load the environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database.
	//db := db.ConnectDB()

	router := gin.New()
	router.Use(gin.Recovery())

	// // Run the application.
	// runApp(db, router)

	// // Close the connection.
	// defer db.Close()
}


// func runApp(db *sql.DB, engine *gin.Engine) {
// 	// Run the application.
// 	router := routes.NewRouter(engine, db)
// 	// Map all routes.
// 	router.MapRoutes()
// 	if err := engine.Run(fmt.Sprintf(":%s", puerto)); err != nil {
// 		panic(err)
// 	}
// }
