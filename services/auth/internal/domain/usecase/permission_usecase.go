package usecase

import (
	"context"
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/domain/repository"
)

type PermissionUseCase interface {
	Create(ctx context.Context, permission *entity.Permission) error
	GetByID(ctx context.Context, id uint) (*entity.Permission, error)
	List(ctx context.Context) ([]*entity.Permission, error)
	Update(ctx context.Context, permission *entity.Permission) error
	Delete(ctx context.Context, id uint) error
}

type permissionUseCase struct {
	permissionRepo repository.PermissionRepository
}

func NewPermissionUseCase(permissionRepo repository.PermissionRepository) PermissionUseCase {
	return &permissionUseCase{
		permissionRepo: permissionRepo,
	}
}

func (u *permissionUseCase) Create(ctx context.Context, permission *entity.Permission) error {
	return u.permissionRepo.Create(permission)
}

func (u *permissionUseCase) GetByID(ctx context.Context, id uint) (*entity.Permission, error) {
	return u.permissionRepo.FindByID(id)
}

func (u *permissionUseCase) List(ctx context.Context) ([]*entity.Permission, error) {
	permissions, _, err := u.permissionRepo.List(1, 100) // Default page 1, limit 100
	if err != nil {
		return nil, err
	}

	// Convert []entity.Permission to []*entity.Permission
	result := make([]*entity.Permission, len(permissions))
	for i := range permissions {
		result[i] = &permissions[i]
	}
	return result, nil
}

func (u *permissionUseCase) Update(ctx context.Context, permission *entity.Permission) error {
	return u.permissionRepo.Update(permission)
}

func (u *permissionUseCase) Delete(ctx context.Context, id uint) error {
	return u.permissionRepo.Delete(id)
}
