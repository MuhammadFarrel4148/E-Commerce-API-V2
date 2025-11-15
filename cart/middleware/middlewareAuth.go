package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "fail",
				"error": "Unauthorized",
			})
			return
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "fail",
				"error": "Unauthorized",
			})
			return
		}

		// Extract the token by removing the "Bearer " prefix
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Store the token in the Gin context for later use in handlers (optional)
		c.Set("bearerToken", token)

		// Proceed to the next handler
		c.Next()
	}
}
