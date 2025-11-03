package main

import (
	"fmt"
	"log"
	"os"
	"product/controller"
	"product/database"
	"product/repository"
	"product/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	db := database.DB

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	router := gin.Default()

	api := router.Group("/product/v1")

	api.POST("/product", productController.CreateProduct)
	api.GET("/product/:id", productController.GetProductByID)
	api.PUT("/product/:id", productController.UpdateProductByID)
	api.DELETE("/product/:id", productController.DeleteProductByID)

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
