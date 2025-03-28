package repository

import (
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/domain/repository"

	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) repository.PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) Create(permission *entity.Permission) error {
	return r.db.Create(permission).Error
}

func (r *permissionRepository) FindByID(id uint) (*entity.Permission, error) {
	var permission entity.Permission
	err := r.db.First(&permission, id).Error
	return &permission, err
}

func (r *permissionRepository) List(page, limit int) ([]entity.Permission, int64, error) {
	var permissions []entity.Permission
	var total int64

	err := r.db.Model(&entity.Permission{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = r.db.Offset(offset).Limit(limit).Find(&permissions).Error
	return permissions, total, err
}

func (r *permissionRepository) Update(permission *entity.Permission) error {
	return r.db.Save(permission).Error
}

func (r *permissionRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Permission{}, id).Error
}

func (r *permissionRepository) FindByAction(action string) ([]entity.Permission, error) {
	var permissions []entity.Permission
	err := r.db.Where("action = ?", action).Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) FindByModule(module string) ([]entity.Permission, error) {
	var permissions []entity.Permission
	err := r.db.Where("module = ?", module).Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) FindByModuleAndAction(module, action string) ([]entity.Permission, error) {
	var permissions []entity.Permission
	err := r.db.Where("module = ? AND action = ?", module, action).Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) FindByName(name string) (*entity.Permission, error) {
	var permission entity.Permission
	err := r.db.Where("name = ?", name).First(&permission).Error
	return &permission, err
}

func (r *permissionRepository) FindByResourceAndAction(resource, action string) (*entity.Permission, error) {
	var permission entity.Permission
	err := r.db.Where("resource = ? AND action = ?", resource, action).First(&permission).Error
	return &permission, err
}

func (r *permissionRepository) FindByRole(roleID uint) ([]entity.Permission, error) {
	var permissions []entity.Permission
	err := r.db.Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id = ?", roleID).
		Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) FindByUser(userID uint) ([]entity.Permission, error) {
	var permissions []entity.Permission
	err := r.db.Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) UpdateStatus(id uint, status entity.PermissionStatus) error {
	return r.db.Model(&entity.Permission{}).Where("id = ?", id).Update("status", status).Error
}

func (r *permissionRepository) UpdateDescription(id uint, description string) error {
	return r.db.Model(&entity.Permission{}).Where("id = ?", id).Update("description", description).Error
}
