package controller

import (
	"net/http"
	"product/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{categoryService}
}

func (h *CategoryController) CreateCategory(c *gin.Context) {
	var input service.InputCategory

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category, err := h.categoryService.CreateCategoryService(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": category,
	})
}

func (h *CategoryController) GetCategoryByID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	product, err := h.categoryService.GetCategoryByID(uint(ID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func (h *CategoryController) UpdateCategoryByID(c *gin.Context) {
	var input service.UpdateCategory

	ID, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category, err := h.categoryService.UpdateCategoryByID(uint(ID), &input)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": category,
	})
}

func (h *CategoryController) DeleteCategoryByID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	category, err := h.categoryService.DeleteCategoryByID(uint(ID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}
