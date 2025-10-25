package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Could not start server", err)
	}
}
