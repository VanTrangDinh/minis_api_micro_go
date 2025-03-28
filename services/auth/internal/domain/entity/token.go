package entity

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID    uint      `gorm:"not null" json:"user_id"`
	Token     string    `gorm:"size:255;not null;unique" json:"token"`
	Type      TokenType `gorm:"size:50;not null" json:"type"`
	ExpiresAt time.Time `json:"expires_at"`
	Revoked   bool      `gorm:"default:false" json:"revoked"`
}

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
	TokenTypeReset   TokenType = "reset"
)

// IsExpired checks if the token has expired
func (t *Token) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}

// IsValid checks if the token is valid (not expired and not revoked)
func (t *Token) IsValid() bool {
	return !t.IsExpired() && !t.Revoked
}

// Revoke marks the token as revoked
func (t *Token) Revoke() {
	t.Revoked = true
}

// TableName specifies the table name for the Token model
func (Token) TableName() string {
	return "tokens"
}
