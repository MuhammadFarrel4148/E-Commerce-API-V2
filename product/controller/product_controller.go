package controller

import (
	"net/http"
	"product/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService}
}

func (h *ProductController) CreateProduct(c *gin.Context) {
	var input service.CreateProductInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := h.productService.CreateProductService(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": product,
	})
}

func (h *ProductController) GetProductByID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	product, err := h.productService.GetProductServiceByID(uint(ID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func (h *ProductController) UpdateProductByID(c *gin.Context) {
	var updateProduct service.UpdateProductInput

	ID, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&updateProduct)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := h.productService.UpdateProductServiceByID(uint(ID), &updateProduct)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": product,
	})
}

func (h *ProductController) DeleteProductByID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	product, err := h.productService.DeleteProductServiceByID(uint(ID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}