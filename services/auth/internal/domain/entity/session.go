package entity

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID    uint      `gorm:"not null" json:"user_id"`
	TokenID   uint      `gorm:"not null" json:"token_id"`
	IP        string    `gorm:"size:45" json:"ip"`
	UserAgent string    `gorm:"size:255" json:"user_agent"`
	ExpiresAt time.Time `json:"expires_at"`
	Active    bool      `gorm:"default:true" json:"active"`
}

// IsExpired checks if the session has expired
func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

// IsValid checks if the session is valid (not expired and active)
func (s *Session) IsValid() bool {
	return !s.IsExpired() && s.Active
}

// Deactivate marks the session as inactive
func (s *Session) Deactivate() {
	s.Active = false
}

// TableName specifies the table name for the Session model
func (Session) TableName() string {
	return "sessions"
}
