package middleware

import (
	"fmt"
	"log"
	"net/http"

	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/domain/repository"
	"minisapi/services/auth/internal/enums"

	"github.com/gin-gonic/gin"
)

// hasPermission checks if a user has the required permission
func hasPermission(user *entity.User, resource enums.PermissionResource, action enums.PermissionAction) bool {
	// Create a map of permissions for faster lookup
	permissionMap := make(map[string]bool)
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			// Create a unique key for the permission
			key := string(permission.Resource) + ":" + string(permission.Action)
			permissionMap[key] = true
		}
	}

	// Check if the required permission exists
	requiredKey := string(resource) + ":" + string(action)
	return permissionMap[requiredKey]
}

// RBACMiddleware checks if the user has the required permission
func RBACMiddleware(userRepo repository.UserRepository, resource, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context (set by AuthMiddleware)
		userID, exists := c.Get("user_id")
		if !exists {
			log.Printf("RBAC: User ID not found in context - Authentication middleware might not be applied correctly")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("Authentication required - Unable to verify user identity for access to %s:%s", resource, action),
				"code":  http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		// Get user with roles and permissions
		user, err := userRepo.FindByID(userID.(uint))
		if err != nil {
			log.Printf("RBAC: Failed to get user with ID %v: %v", userID, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "User information could not be retrieved - Please try again later",
				"code":  http.StatusInternalServerError,
			})
			c.Abort()
			return
		}

		// Convert string to enum
		resourceEnum := enums.PermissionResource(resource)
		actionEnum := enums.PermissionAction(action)

		// Check if user has the required permission
		if !hasPermission(user, resourceEnum, actionEnum) {
			log.Printf("RBAC: User %v does not have permission for %s:%s", userID, resource, action)
			c.JSON(http.StatusForbidden, gin.H{
				"error": fmt.Sprintf("Access denied - You don't have permission to %s on %s resource", action, resource),
				"code":  http.StatusForbidden,
			})
			c.Abort()
			return
		}

		log.Printf("RBAC: User %v has permission for %s:%s", userID, resource, action)
		c.Next()
	}
}
