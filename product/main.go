package main

import (
	"fmt"
	"log"
	"os"
	"product/caching"
	"product/controller"
	"product/database"
	"product/repository"
	"product/service"

	"github.com/gin-gonic/gin"
)

func main() {
	caching.ConnectRedis()

	database.ConnectDB()
	db := database.DB

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	inventoryRepository := repository.NewInventoryRepository(db)
	inventoryService := service.NewInventoryService(inventoryRepository)
	inventoryController := controller.NewInventoryController(inventoryService)

	router := gin.Default()

	apiProduct := router.Group("/product/v1")
	apiCategory := router.Group("/category/v1")
	apiInventory := router.Group("/inventory/v1")

	apiProduct.POST("/product", productController.CreateProduct)
	apiProduct.GET("/product/:id", productController.GetProductByID)
	apiProduct.PUT("/product/:id", productController.UpdateProductByID)
	apiProduct.DELETE("/product/:id", productController.DeleteProductByID)

	apiCategory.POST("/category", categoryController.CreateCategory)
	apiCategory.GET("/category/:id", categoryController.GetCategoryByID)
	apiCategory.PUT("/category/:id", categoryController.UpdateCategoryByID)
	apiCategory.DELETE("/category/:id", categoryController.DeleteCategoryByID)

	apiInventory.POST("/inventory", inventoryController.CreateInventory)
	apiInventory.GET("/inventory/:id", inventoryController.GetInventoryByID)
	apiInventory.PUT("/inventory/:id", inventoryController.UpdateInventoryByID)
	apiInventory.DELETE("/inventory/:id", inventoryController.DeleteInventoryByID)

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
