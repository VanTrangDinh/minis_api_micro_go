package repository

import (
	"context"
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/domain/repository"
	"minisapi/services/auth/internal/pkg/errors"
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *userRepository) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.ErrDatabase
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.ErrDatabase
	}
	return &user, nil
}

func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email).Count(&count)
	if result.Error != nil {
		return false, errors.ErrDatabase
	}
	return count > 0, nil
}

func (r *userRepository) UpdateEmailVerified(ctx context.Context, id uint, verified bool) error {
	result := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Update("email_verified", verified)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *userRepository) UpdatePassword(id uint, password string) error {
	result := r.db.Model(&entity.User{}).Where("id = ?", id).Update("password", password)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *userRepository) UpdateTwoFactor(ctx context.Context, id uint, enabled bool, secret string) error {
	result := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"two_factor_enabled": enabled,
		"two_factor_secret":  secret,
	})
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *userRepository) Delete(id uint) error {
	result := r.db.Delete(&entity.User{}, id)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *userRepository) FindByPermission(permissionID uint) ([]entity.User, error) {
	var users []entity.User
	result := r.db.Joins("JOIN role_permissions ON roles.id = role_permissions.role_id").
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Where("role_permissions.permission_id = ?", permissionID).
		Find(&users)
	if result.Error != nil {
		return nil, errors.ErrDatabase
	}
	return users, nil
}

func (r *userRepository) FindByRole(roleID uint) ([]entity.User, error) {
	var users []entity.User
	result := r.db.Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Where("user_roles.role_id = ?", roleID).
		Find(&users)
	if result.Error != nil {
		return nil, errors.ErrDatabase
	}
	return users, nil
}

func (r *userRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.ErrDatabase
	}
	return &user, nil
}

func (r *userRepository) IncrementFailedLogins(id uint) error {
	result := r.db.Model(&entity.User{}).Where("id = ?", id).
		UpdateColumn("failed_login_attempts", gorm.Expr("failed_login_attempts + ?", 1))
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *userRepository) List(page, limit int) ([]entity.User, int64, error) {
	var users []entity.User
	var count int64

	// Get total count
	if err := r.db.Model(&entity.User{}).Count(&count).Error; err != nil {
		return nil, 0, errors.ErrDatabase
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get paginated results
	result := r.db.Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, 0, errors.ErrDatabase
	}

	return users, count, nil
}

func (r *userRepository) LockAccount(id uint, duration int) error {
	// Calculate unlock time based on duration in minutes
	unlockTime := time.Now().Add(time.Duration(duration) * time.Minute)

	result := r.db.Model(&entity.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"locked":       true,
		"locked_until": unlockTime,
	})
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *userRepository) UnlockAccount(id uint) error {
	result := r.db.Model(&entity.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"locked":       false,
		"locked_until": nil,
	})
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *userRepository) ResetFailedLogins(id uint) error {
	result := r.db.Model(&entity.User{}).Where("id = ?", id).Update("failed_login_attempts", 0)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *userRepository) UpdateLastLogin(id uint, ip string) error {
	result := r.db.Model(&entity.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"last_login_at": time.Now(),
		"last_login_ip": ip,
	})
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *userRepository) UpdateStatus(id uint, status entity.UserStatus) error {
	result := r.db.Model(&entity.User{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *userRepository) Update(user *entity.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}
