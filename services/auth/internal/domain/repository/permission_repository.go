package repository

import (
	"minisapi/services/auth/internal/domain/entity"
)

type PermissionRepository interface {
	Create(permission *entity.Permission) error
	FindByID(id uint) (*entity.Permission, error)
	FindByName(name string) (*entity.Permission, error)
	FindByResourceAndAction(resource, action string) (*entity.Permission, error)
	Update(permission *entity.Permission) error
	Delete(id uint) error
	List(page, limit int) ([]entity.Permission, int64, error)
	FindByRole(roleID uint) ([]entity.Permission, error)
	FindByUser(userID uint) ([]entity.Permission, error)
	UpdateStatus(id uint, status entity.PermissionStatus) error
	UpdateDescription(id uint, description string) error
	FindByModule(module string) ([]entity.Permission, error)
	FindByAction(action string) ([]entity.Permission, error)
	FindByModuleAndAction(module, action string) ([]entity.Permission, error)
}
