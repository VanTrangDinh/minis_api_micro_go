package repository

import (
	"minisapi/services/auth/internal/domain/entity"
)

type RoleRepository interface {
	Create(role *entity.Role) error
	Update(role *entity.Role) error
	Delete(id uint) error
	FindByID(id uint) (*entity.Role, error)
	FindByName(name string) (*entity.Role, error)
	List(page, limit int) ([]entity.Role, int64, error)
	AssignPermission(roleID uint, permissionID uint) error
	RemovePermission(roleID uint, permissionID uint) error
	FindByPermission(permissionID uint) ([]entity.Role, error)
	UpdateStatus(id uint, status entity.RoleStatus) error
	UpdateDescription(id uint, description string) error
	FindByUser(userID uint) ([]entity.Role, error)
}
