package middleware

import (
	"strings"

	"minisapi/services/notification/internal/response"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates the JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.SendError(c, response.Unauthorized, "authorization header is required")
			c.Abort()
			return
		}

		// Check Bearer token format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.SendError(c, response.Unauthorized, "invalid authorization header format")
			c.Abort()
			return
		}

		// TODO: Validate JWT token
		// For now, we just pass the token to the next handler
		c.Set("token", parts[1])
		c.Next()
	}
}
