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

type InventoryController struct {
	inventoryService service.InventoryService
}

func NewInventoryController(inventoryService service.InventoryService) *InventoryController {
	return &InventoryController{inventoryService}
}

func (h *InventoryController) CreateInventory(c *gin.Context) {
	var input model.InputInventory
	var validationErr *exceptions.ErrValidation

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error": err.Error(),
		})
		return
	}

	inventory, err := h.inventoryService.CreateInventory(input)
	if err != nil {
		if errors.Is(err, exceptions.ErrProductIDFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "fail",
				"error": err.Error(),
			})
			return
		} else if errors.Is(err, validationErr) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "fail",
				"error": validationErr.Details,
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "fail",
				"error": "terjadi kesalahan pada server kami",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": inventory,
	})
}

func (h *InventoryController) GetInventoryByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error": err.Error(),
		})
		return
	}

	inventory, err := h.inventoryService.GetInventoryByID(uint(ID))
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
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": inventory,
	})
}

func (h *InventoryController) UpdateInventoryByID(c *gin.Context) {
	var input model.UpdateInventory
	var validationErr *exceptions.ErrValidation

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error": err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error": err.Error(),
		})
		return
	}

	inventory, err := h.inventoryService.UpdateInventoryByID(uint(ID), &input)
	if err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "fail",
				"error": err.Error(),
			})
			return
		} else if errors.Is(err, validationErr) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "fail",
				"error": validationErr.Details,
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "fail",
				"error": err.Error(),
			})
			return
		}	
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": inventory,
	})
}

func (h *InventoryController) DeleteInventoryByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error": err.Error(),
		})
	}

	inventory, err := h.inventoryService.DeleteInventoryByID(uint(ID))
	if err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "fail",
				"error": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "fail",
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": inventory,
	})
}
