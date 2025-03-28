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

func (r *permissionRepository) List() ([]*entity.Permission, error) {
	var permissions []*entity.Permission
	err := r.db.Find(&permissions).Error
	return permissions, err
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
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
