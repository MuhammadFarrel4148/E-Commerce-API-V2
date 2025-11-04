package controller

import (
	"net/http"
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
	var input service.InputInventory

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	inventory, err := h.inventoryService.CreateInventory(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": inventory,
	})
}

func (h *InventoryController) GetInventoryByID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	inventory, err := h.inventoryService.GetInventoryByID(uint(ID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": inventory,
	})
}

func (h *InventoryController) UpdateInventoryByID(c *gin.Context) {
	var input service.UpdateInventory

	ID, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	inventory, err := h.inventoryService.UpdateInventoryByID(uint(ID), &input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": inventory,
	})
}

func (h *InventoryController) DeleteInventoryByID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	inventory, err := h.inventoryService.DeleteInventoryByID(uint(ID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": inventory,
	})
}
