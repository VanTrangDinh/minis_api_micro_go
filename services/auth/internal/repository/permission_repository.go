package repository

import (
	stderrors "errors"
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/pkg/errors"

	"gorm.io/gorm"
)

// PermissionRepository defines the interface for permission database operations
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

// permissionRepository implements PermissionRepository interface
type permissionRepository struct {
	db *gorm.DB
}

// NewPermissionRepository creates a new instance of PermissionRepository
func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{
		db: db,
	}
}

// Create creates a new permission
func (r *permissionRepository) Create(permission *entity.Permission) error {
	if err := r.db.Create(permission).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}

// FindByID finds a permission by ID
func (r *permissionRepository) FindByID(id uint) (*entity.Permission, error) {
	var permission entity.Permission
	if err := r.db.First(&permission, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrPermissionNotFound
		}
		return nil, errors.ErrDatabase
	}
	return &permission, nil
}

// FindByName finds a permission by name
func (r *permissionRepository) FindByName(name string) (*entity.Permission, error) {
	var permission entity.Permission
	if err := r.db.Where("name = ?", name).First(&permission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrPermissionNotFound
		}
		return nil, errors.ErrDatabase
	}
	return &permission, nil
}

// FindByResourceAndAction finds a permission by resource and action
func (r *permissionRepository) FindByResourceAndAction(resource, action string) (*entity.Permission, error) {
	var permission entity.Permission
	err := r.db.Preload("Roles").Where("resource = ? AND action = ?", resource, action).First(&permission).Error
	if err != nil {
		if stderrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &permission, nil
}

// Update updates a permission
func (r *permissionRepository) Update(permission *entity.Permission) error {
	if err := r.db.Save(permission).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}

// Delete deletes a permission
func (r *permissionRepository) Delete(id uint) error {
	if err := r.db.Delete(&entity.Permission{}, id).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}

// List returns a list of permissions
func (r *permissionRepository) List(page, limit int) ([]entity.Permission, int64, error) {
	var permissions []entity.Permission
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Model(&entity.Permission{}).Count(&total).Error; err != nil {
		return nil, 0, errors.ErrDatabase
	}

	if err := r.db.Offset(offset).Limit(limit).Find(&permissions).Error; err != nil {
		return nil, 0, errors.ErrDatabase
	}

	return permissions, total, nil
}

func (r *permissionRepository) FindByRole(roleID uint) ([]entity.Permission, error) {
	var permissions []entity.Permission
	if err := r.db.Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id = ?", roleID).
		Find(&permissions).Error; err != nil {
		return nil, errors.ErrDatabase
	}
	return permissions, nil
}

func (r *permissionRepository) FindByUser(userID uint) ([]entity.Permission, error) {
	var permissions []entity.Permission
	if err := r.db.Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Find(&permissions).Error; err != nil {
		return nil, errors.ErrDatabase
	}
	return permissions, nil
}

func (r *permissionRepository) UpdateStatus(id uint, status entity.PermissionStatus) error {
	if err := r.db.Model(&entity.Permission{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *permissionRepository) UpdateDescription(id uint, description string) error {
	if err := r.db.Model(&entity.Permission{}).Where("id = ?", id).Update("description", description).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *permissionRepository) FindByModule(module string) ([]entity.Permission, error) {
	var permissions []entity.Permission
	if err := r.db.Where("module = ?", module).Find(&permissions).Error; err != nil {
		return nil, errors.ErrDatabase
	}
	return permissions, nil
}

func (r *permissionRepository) FindByAction(action string) ([]entity.Permission, error) {
	var permissions []entity.Permission
	if err := r.db.Where("action = ?", action).Find(&permissions).Error; err != nil {
		return nil, errors.ErrDatabase
	}
	return permissions, nil
}

func (r *permissionRepository) FindByModuleAndAction(module, action string) ([]entity.Permission, error) {
	var permissions []entity.Permission
	if err := r.db.Where("module = ? AND action = ?", module, action).Find(&permissions).Error; err != nil {
		return nil, errors.ErrDatabase
	}
	return permissions, nil
}
