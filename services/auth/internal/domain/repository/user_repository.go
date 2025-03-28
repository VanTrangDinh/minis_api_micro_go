package repository

import (
	"context"
	"minisapi/services/auth/internal/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id uint) error
	FindByID(id uint) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
	List(page, limit int) ([]entity.User, int64, error)
	FindByRole(roleID uint) ([]entity.User, error)
	FindByPermission(permissionID uint) ([]entity.User, error)
	UpdateStatus(id uint, status entity.UserStatus) error
	UpdatePassword(id uint, hashedPassword string) error
	UpdateLastLogin(id uint, ip string) error
	IncrementFailedLogins(id uint) error
	ResetFailedLogins(id uint) error
	LockAccount(id uint, duration int) error
	UnlockAccount(id uint) error
	UpdateTwoFactor(ctx context.Context, id uint, enabled bool, secret string) error
}
