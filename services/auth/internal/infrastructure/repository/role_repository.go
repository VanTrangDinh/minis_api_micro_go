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

func (r *roleRepository) FindByID(id uint) (*entity.Role, error) {
	var role entity.Role
	err := r.db.First(&role, id).Error
	return &role, err
}

func (r *roleRepository) FindByName(name string) (*entity.Role, error) {
	var role entity.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	return &role, err
}

func (r *roleRepository) List(page, limit int) ([]entity.Role, int64, error) {
	var roles []entity.Role
	var total int64

	err := r.db.Model(&entity.Role{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = r.db.Offset(offset).Limit(limit).Find(&roles).Error
	return roles, total, err
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

func (r *roleRepository) RemovePermission(roleID uint, permissionID uint) error {
	return r.db.Model(&entity.Role{ID: roleID}).Association("Permissions").Delete(&entity.Permission{ID: permissionID})
}

func (r *roleRepository) FindByPermission(permissionID uint) ([]entity.Role, error) {
	var roles []entity.Role
	err := r.db.Joins("JOIN role_permissions ON roles.id = role_permissions.role_id").
		Where("role_permissions.permission_id = ?", permissionID).
		Find(&roles).Error
	return roles, err
}

func (r *roleRepository) UpdateStatus(id uint, status entity.RoleStatus) error {
	return r.db.Model(&entity.Role{}).Where("id = ?", id).Update("status", status).Error
}

func (r *roleRepository) UpdateDescription(id uint, description string) error {
	return r.db.Model(&entity.Role{}).Where("id = ?", id).Update("description", description).Error
}

func (r *roleRepository) FindByUser(userID uint) ([]entity.Role, error) {
	var roles []entity.Role
	err := r.db.Joins("JOIN user_roles ON roles.id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}
