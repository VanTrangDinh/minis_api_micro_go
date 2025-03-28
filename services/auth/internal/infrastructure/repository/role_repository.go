package repository

import (
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/domain/repository"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *entity.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetByID(id uint) (*entity.Role, error) {
	var role entity.Role
	err := r.db.First(&role, id).Error
	return &role, err
}

func (r *roleRepository) List() ([]*entity.Role, error) {
	var roles []*entity.Role
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *roleRepository) Update(role *entity.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Role{}, id).Error
}

func (r *roleRepository) AssignPermission(roleID uint, permissionID uint) error {
	return r.db.Model(&entity.Role{ID: roleID}).Association("Permissions").Append(&entity.Permission{ID: permissionID})
}
