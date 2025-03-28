package middleware

import (
	"minisapi/services/auth/internal/domain/service"
	"minisapi/services/auth/internal/pkg/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authService service.AuthService
}

func NewAuthMiddleware(authService service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUnauthorized.Error()})
			c.Abort()
			return
		}

		// Extract token from Bearer header
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUnauthorized.Error()})
			c.Abort()
			return
		}

		token := parts[1]
		userID, err := m.authService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUnauthorized.Error()})
			c.Abort()
			return
		}

		// Set user ID in context
		c.Set("user_id", userID)
		c.Next()
	}
}

func (m *AuthMiddleware) RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		userRoles, err := m.authService.GetUserRoles(userID)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": errors.ErrForbidden.Error()})
			c.Abort()
			return
		}

		// Check if user has any of the required roles
		hasRole := false
		for _, role := range roles {
			for _, userRole := range userRoles {
				if role == userRole {
					hasRole = true
					break
				}
			}
			if hasRole {
				break
			}
		}

		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{"error": errors.ErrForbidden.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (m *AuthMiddleware) RequirePermission(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		userPermissions, err := m.authService.GetUserPermissions(userID)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": errors.ErrForbidden.Error()})
			c.Abort()
			return
		}

		// Check if user has all required permissions
		for _, permission := range permissions {
			hasPermission := false
			for _, userPermission := range userPermissions {
				if permission == userPermission {
					hasPermission = true
					break
				}
			}
			if !hasPermission {
				c.JSON(http.StatusForbidden, gin.H{"error": errors.ErrForbidden.Error()})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

func (m *AuthMiddleware) RequireTwoFactor() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		enabled, err := m.authService.IsTwoFactorEnabled(userID)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": errors.ErrForbidden.Error()})
			c.Abort()
			return
		}

		if !enabled {
			c.JSON(http.StatusForbidden, gin.H{"error": errors.ErrTwoFactorRequired.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
