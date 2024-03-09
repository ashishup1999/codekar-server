package main

import (
	"codekar/app/db"
	"codekar/app/handlers"
	"log"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables from .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Println("Error loading .env file:", err)
	// 	return
	// }

	//setup db
	dbURI := url.QueryEscape(os.Getenv("DB_URI"))
	db.ConnectToMongoDB(dbURI)

	//api routes setup
	r := gin.Default()
	handlers.SetUphandler(r)

	//start server
	// if err := r.Run(":" + os.Getenv("PORT")); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
