package main

import (
	"codekar/app/db"
	"codekar/app/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load(".env")

	//setup db
	dbURI := os.Getenv("DB_URI")
	db.ConnectToMongoDB(dbURI)

	//create a route handler from gin
	r := gin.Default()

	//set api routes
	handlers.SetUphandler(r)

	// start server
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
