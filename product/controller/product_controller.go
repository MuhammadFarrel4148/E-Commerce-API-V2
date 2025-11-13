package controller

import (
	"errors"
	"net/http"
	"product/exceptions"
	"product/model"
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
	var inputProduct model.CreateProductInput
	var validationErr *exceptions.ErrValidation

	if err := c.ShouldBindJSON(&inputProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	product, err := h.productService.CreateProductService(c.Request.Context(), inputProduct)

	if err != nil {
		if errors.As(err, &validationErr) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "fail",
				"error":  validationErr.Details,
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "fail",
				"error":  "terjadi kesalahan pada server kami",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   product,
	})
}

func (h *ProductController) GetProductByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "invalid product ID",
		})
		return
	}

	product, err := h.productService.GetProductServiceByID(c.Request.Context(), uint(ID))

	if err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "fail",
				"error":  err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "fail",
				"error":  "terjadi kesalahan pada server kami",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}

func (h *ProductController) UpdateProductByID(c *gin.Context) {
	var updateProduct model.UpdateProductInput

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "invalid product ID",
		})
		return
	}

	err = c.ShouldBindJSON(&updateProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	product, err := h.productService.UpdateProductServiceByID(c.Request.Context(), uint(ID), updateProduct)
	if err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "fail",
				"error":  err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "fail",
				"error":  "terjadi kesalahan pada server kami",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}

func (h *ProductController) DeleteProductByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "invalid product ID",
		})
		return
	}

	product, err := h.productService.DeleteProductServiceByID(c.Request.Context(), uint(ID))
	if err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "fail",
				"error": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "fail",
				"error":  "terjadi kesalahan pada server kami",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}
