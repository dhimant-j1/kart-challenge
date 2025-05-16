package middleware

import (
	"backend-challenge/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIKeyAuth middleware checks for a valid API key in the header
func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if endpoint requires authentication
		if c.Request.URL.Path == "/api/order" && c.Request.Method == "POST" {
			apiKey := c.GetHeader("api_key")
			if apiKey != "apitest" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, models.APIResponse{
					Code:    http.StatusUnauthorized,
					Type:    "error",
					Message: "Invalid or missing API key",
				})
				return
			}
		}

		c.Next()
	}
}

// CORS middleware adds CORS headers to responses
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, api_key, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
