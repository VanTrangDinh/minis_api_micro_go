package entity

import (
	"time"

	"gorm.io/gorm"
)

type PermissionStatus string

const (
	PermissionStatusActive   PermissionStatus = "active"
	PermissionStatusInactive PermissionStatus = "inactive"
	PermissionStatusDeleted  PermissionStatus = "deleted"
)

type Permission struct {
	ID        uint             `gorm:"primarykey" json:"id"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
	Status    PermissionStatus `gorm:"type:varchar(20);default:'active'" json:"status"`

	Name        string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"size:255" json:"description"`
	Resource    string `gorm:"size:50;not null" json:"resource"` // e.g., "users", "posts"
	Action      string `gorm:"size:50;not null" json:"action"`   // e.g., "create", "read", "update", "delete"

	Roles []Role `gorm:"many2many:role_permissions;" json:"roles"`
}

// BeforeSave is a hook that is called before saving the permission
func (p *Permission) BeforeSave(tx *gorm.DB) error {
	// Set name based on resource and action if not set
	if p.Name == "" {
		p.Name = p.Resource + ":" + p.Action
	}
	return nil
}

// TableName specifies the table name for the Permission model
func (Permission) TableName() string {
	return "permissions"
}
