package main

import (
	"codekar/app/db"
	"codekar/app/handlers"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

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
