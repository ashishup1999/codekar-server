package main

import (
	"codekar/app/db"
	"codekar/app/handlers"
	"codekar/app/services"
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
	dbURI := fmt.Sprintf("mongodb+srv://%s:%s@%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"))
    fmt.Println(dbURI)
	dbClient, err := db.ConnectToMongoDB(dbURI)
	if err != nil {
		fmt.Println(err.Error())
	}

	//passing db client to services
	services.DbClient = dbClient

	//api routes setup
	r := gin.Default()
	handlers.SetUphandler(r)

	//start server
	if err := r.Run(":"+os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
