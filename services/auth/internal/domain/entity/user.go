package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Email            string     `gorm:"size:255;not null;unique" json:"email"`
	Username         string     `gorm:"size:255;not null;unique" json:"username"`
	Password         string     `gorm:"size:255;not null" json:"-"`
	FirstName        string     `gorm:"size:255" json:"first_name"`
	LastName         string     `gorm:"size:255" json:"last_name"`
	Phone            string     `json:"phone"`
	Active           bool       `json:"active" gorm:"default:true"`
	EmailVerified    bool       `json:"email_verified" gorm:"default:false"`
	TwoFactorEnabled bool       `json:"two_factor_enabled" gorm:"default:false"`
	LastLoginAt      *time.Time `json:"last_login_at"`
	LastLoginIP      string     `json:"last_login_ip"`
	FailedLogins     int        `json:"failed_logins" gorm:"default:0"`
	LockedUntil      *time.Time `json:"locked_until"`

	Status UserStatus `gorm:"type:varchar(20);default:'active'" json:"status"`
	Type   UserType   `gorm:"type:varchar(20);default:'user'" json:"type"`

	Roles []Role `gorm:"many2many:user_roles;" json:"roles"`
}

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusLocked   UserStatus = "locked"
	UserStatusDeleted  UserStatus = "deleted"
)

type UserType string

const (
	UserTypeAdmin UserType = "admin"
	UserTypeUser  UserType = "user"
)

// DefaultRole returns the default role for a user type
func (ut UserType) DefaultRole() RoleType {
	switch ut {
	case UserTypeAdmin:
		return RoleTypeAdmin
	default:
		return RoleTypeUser
	}
}

// BeforeSave is a hook that is called before saving the user
func (u *User) BeforeSave(tx *gorm.DB) error {
	// Set default role based on user type if no roles are assigned
	if len(u.Roles) == 0 {
		defaultRole := u.Type.DefaultRole()
		u.Roles = []Role{{Name: string(defaultRole)}}
	}
	return nil
}

// IsLocked checks if the user account is locked
func (u *User) IsLocked() bool {
	if u.LockedUntil == nil {
		return false
	}
	return time.Now().Before(*u.LockedUntil)
}

// LockAccount locks the user account for a specified duration
func (u *User) LockAccount(duration time.Duration) {
	lockedUntil := time.Now().Add(duration)
	u.LockedUntil = &lockedUntil
	u.Status = UserStatusLocked
}

// UnlockAccount unlocks the user account
func (u *User) UnlockAccount() {
	u.LockedUntil = nil
	u.Status = UserStatusActive
	u.FailedLogins = 0
}

// IncrementFailedLogins increments the failed login counter
func (u *User) IncrementFailedLogins() {
	u.FailedLogins++
	if u.FailedLogins >= 5 {
		u.LockAccount(30 * time.Minute)
	}
}

// ResetFailedLogins resets the failed login counter
func (u *User) ResetFailedLogins() {
	u.FailedLogins = 0
}

// UpdateLastLogin updates the last login information
func (u *User) UpdateLastLogin(ip string) {
	now := time.Now()
	u.LastLoginAt = &now
	u.LastLoginIP = ip
	u.ResetFailedLogins()
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}
