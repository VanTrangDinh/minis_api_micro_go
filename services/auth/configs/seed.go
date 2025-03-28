package configs

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/domain/enums"
	"minisapi/services/auth/internal/pkg/utils"
)

// SeedData initializes the database with default data
func SeedData(db *gorm.DB) error {
	// Create permissions first
	if err := seedPermissions(db); err != nil {
		return fmt.Errorf("failed to seed permissions: %v", err)
	}

	// Create admin role
	adminRole, err := seedAdminRole(db)
	if err != nil {
		return fmt.Errorf("failed to seed admin role: %v", err)
	}

	// Create admin user
	if err := seedAdminUser(db, adminRole); err != nil {
		return fmt.Errorf("failed to seed admin user: %v", err)
	}

	return nil
}

// seedPermissions creates all necessary permissions
func seedPermissions(db *gorm.DB) error {
	// Resources
	resources := []enums.PermissionResource{
		enums.ResourceUser,
		enums.ResourceRole,
		enums.ResourcePermission,
		enums.ResourceAuth,
		enums.ResourceToken,
	}

	// Actions
	actions := []enums.PermissionAction{
		enums.ActionCreate,
		enums.ActionRead,
		enums.ActionUpdate,
		enums.ActionDelete,
		enums.ActionAssign,
		enums.ActionRemove,
		enums.ActionActivate,
		enums.ActionDeactivate,
	}

	// Create combinations of resources and actions
	for _, resource := range resources {
		for _, action := range actions {
			// Skip combinations that don't make sense
			if !isValidPermission(resource, action) {
				continue
			}

			permName := string(resource) + ":" + string(action)
			var permission entity.Permission
			if err := db.Where("name = ?", permName).First(&permission).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// Create permission
					permission = entity.Permission{
						Name:        permName,
						Description: fmt.Sprintf("Permission to %s %s", action, resource),
						Resource:    string(resource),
						Action:      string(action),
					}

					if err := db.Create(&permission).Error; err != nil {
						return err
					}
					log.Printf("Created permission: %s", permName)
				} else {
					return err
				}
			}
		}
	}

	return nil
}

// isValidPermission checks if a resource-action combination makes sense
func isValidPermission(resource enums.PermissionResource, action enums.PermissionAction) bool {
	// Some actions only apply to certain resources
	if action == enums.ActionAssign || action == enums.ActionRemove {
		return resource == enums.ResourceRole || resource == enums.ResourcePermission
	}

	if action == enums.ActionActivate || action == enums.ActionDeactivate {
		return resource == enums.ResourceUser
	}

	// Auth resource has specific actions
	if resource == enums.ResourceAuth {
		return action == enums.ActionCreate || action == enums.ActionRead
	}

	// Token resource has specific actions
	if resource == enums.ResourceToken {
		return action == enums.ActionCreate || action == enums.ActionRead
	}

	// Default actions (CRUD) apply to most resources
	return true
}

// seedAdminRole creates admin role with all permissions
func seedAdminRole(db *gorm.DB) (entity.Role, error) {
	var adminRole entity.Role
	if err := db.Where("name = ?", string(entity.RoleTypeAdmin)).First(&adminRole).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create admin role
			adminRole = entity.Role{
				Name:        string(entity.RoleTypeAdmin),
				Description: "Administrator role with full system access",
				Type:        entity.RoleTypeAdmin,
			}

			// Get all permissions
			var permissions []entity.Permission
			if err := db.Find(&permissions).Error; err != nil {
				return adminRole, err
			}

			// Assign all permissions to admin role
			adminRole.Permissions = permissions

			// Save admin role
			if err := db.Create(&adminRole).Error; err != nil {
				return adminRole, err
			}
			log.Printf("Admin role created successfully with %d permissions", len(permissions))
		} else {
			return adminRole, err
		}
	} else {
		// If role exists, ensure it has all permissions
		var permissions []entity.Permission
		if err := db.Find(&permissions).Error; err != nil {
			return adminRole, err
		}

		// Update permissions
		if err := db.Model(&adminRole).Association("Permissions").Replace(permissions); err != nil {
			return adminRole, err
		}
		log.Printf("Updated admin role with %d permissions", len(permissions))
	}

	// Preload permissions to ensure they are loaded
	if err := db.Preload("Permissions").First(&adminRole, adminRole.ID).Error; err != nil {
		return adminRole, err
	}

	return adminRole, nil
}

// seedAdminUser creates admin user with admin role
func seedAdminUser(db *gorm.DB, adminRole entity.Role) error {
	var adminUser entity.User
	if err := db.Where("username = ?", "admin").First(&adminUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Hash password
			hashedPassword, err := utils.HashPassword("admin123")
			if err != nil {
				return err
			}

			// Create admin user
			adminUser = entity.User{
				Username:  "admin",
				Email:     "admin@example.com",
				Password:  hashedPassword,
				FirstName: "Admin",
				LastName:  "User",
				Active:    true,
				Status:    entity.UserStatusActive,
				Type:      entity.UserTypeAdmin,
				Roles:     []entity.Role{adminRole},
			}

			if err := db.Create(&adminUser).Error; err != nil {
				return err
			}
			log.Printf("Admin user created successfully")
		} else {
			return err
		}
	}

	return nil
}
