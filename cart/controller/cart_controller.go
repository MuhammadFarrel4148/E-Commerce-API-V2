package controller

import (
	"cart/exceptions"
	"cart/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cartController struct {
	cartService service.CartService
}

func NewCartController(cartService service.CartService) *cartController {
	return &cartController{cartService}
}

func (h *cartController) CreateCartController(c *gin.Context) {
	userID, exists := c.Get("bearerToken")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "userID not found",
		})
		return
	}
	userId, _ := userID.(string)

	cart, err := service.CreateCartService(userId)
	if err != nil {
		if errors.Is(err, exceptions.ErrCartFound) {
			c.JSON(http.StatusBadRequest, gin.H{
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

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": cart
	})
}
