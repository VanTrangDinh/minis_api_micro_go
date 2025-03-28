package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string     `gorm:"size:255;not null;unique" json:"name"`
	Description string     `gorm:"size:255" json:"description"`
	Type        RoleType   `gorm:"type:varchar(20);default:'user'" json:"type"`
	Status      RoleStatus `gorm:"type:varchar(20);default:'active'" json:"status"`

	Users       []User       `gorm:"many2many:user_roles;" json:"users"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type RoleType string

const (
	RoleTypeAdmin RoleType = "admin"
	RoleTypeUser  RoleType = "user"
)

type RoleStatus string

const (
	RoleStatusActive   RoleStatus = "active"
	RoleStatusInactive RoleStatus = "inactive"
	RoleStatusDeleted  RoleStatus = "deleted"
)

// DefaultRole returns the default role for a user type
func (rt RoleType) DefaultRole() RoleType {
	switch rt {
	case RoleTypeAdmin:
		return RoleTypeAdmin
	default:
		return RoleTypeUser
	}
}

// DefaultPermissions returns the default permissions for a role type
func (rt RoleType) DefaultPermissions() []Permission {
	switch rt {
	case RoleTypeAdmin:
		return []Permission{
			{Resource: "*", Action: "*"},
		}
	default:
		return []Permission{
			{Resource: "profile", Action: "read"},
			{Resource: "profile", Action: "update"},
		}
	}
}

// BeforeSave is a hook that is called before saving the role
func (r *Role) BeforeSave(tx *gorm.DB) error {
	// Set default permissions based on role type if no permissions are assigned
	if len(r.Permissions) == 0 {
		r.Permissions = make([]Permission, 0)
		for _, p := range r.Type.DefaultPermissions() {
			r.Permissions = append(r.Permissions, Permission{
				Resource: string(p.Resource),
				Action:   string(p.Action),
			})
		}
	}
	return nil
}

// TableName specifies the table name for the Role model
func (Role) TableName() string {
	return "roles"
}
