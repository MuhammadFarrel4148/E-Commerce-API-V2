package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"product/caching"
	"product/exceptions"
	"product/model"
	"product/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{categoryService}
}

func (h *CategoryController) CreateCategory(c *gin.Context) {
	var input model.InputCategory

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	category, err := h.categoryService.CreateCategoryService(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	key := fmt.Sprintf("category-%d", category.CategoryID)
	err = caching.Set(key, category, 15*time.Minute)

	if err != nil {
		log.Printf("WARNING: Gagal menyimpan cache untuk key %s. Error: %v", key, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   category,
	})
}

func (h *CategoryController) GetCategoryByID(c *gin.Context) {
	var categoryOutput model.CategoryOutput

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	key := fmt.Sprintf("category-%d", ID)
	category, err := caching.Get(key)

	if err == nil && category != "" {
		err = json.Unmarshal([]byte(category), &categoryOutput)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data":   categoryOutput,
			})
			return
		}

		log.Printf("ERROR: Data cache rusak untuk key %s: %v", key, err)
	}

	categoryFromDB, err := h.categoryService.GetCategoryByID(uint(ID))
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

	err = caching.Set(key, categoryFromDB, 15*time.Minute)
	if err != nil {
		log.Printf("WARNING: Gagal set cache untuk key %s: %v", key, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   categoryFromDB,
	})
}

func (h *CategoryController) UpdateCategoryByID(c *gin.Context) {
	var input model.UpdateCategory

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	category, err := h.categoryService.UpdateCategoryByID(uint(ID), &input)
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
				"error":  err.Error(),
			})
			return
		}
	}

	key := fmt.Sprintf("category-%d", ID)
	err = caching.Set(key, category, 15*time.Minute)
	if err != nil {
		log.Printf("WARNING: Gagal set cache untuk key %s: %v", key, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   category,
	})
}

func (h *CategoryController) DeleteCategoryByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	category, err := h.categoryService.DeleteCategoryByID(uint(ID))
	if err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "fail",
				"error":  err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "fail",
				"error":  err.Error(),
			})
			return
		}
	}

	key := fmt.Sprintf("category-%d", ID)
	err = caching.Del(key)
	if err != nil {
		log.Printf("WARNING: Gagal set cache untuk key %s: %v", key, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   category,
	})
}
