package main

import (
	"cart/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	database.ConnectDatabase()
	// db := database.DB

	router := gin.Default()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Port := os.Getenv("SERVER_PORT")
	log.Printf("server starting on port %s", Port)

	if err := router.Run(Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
