package usecase

import (
	"context"
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/domain/repository"
)

type RoleUseCase interface {
	Create(ctx context.Context, role *entity.Role) error
	GetByID(ctx context.Context, id uint) (*entity.Role, error)
	List(ctx context.Context) ([]*entity.Role, error)
	Update(ctx context.Context, role *entity.Role) error
	Delete(ctx context.Context, id uint) error
	AssignPermission(ctx context.Context, roleID uint, permissionID uint) error
}

type roleUseCase struct {
	roleRepo       repository.RoleRepository
	permissionRepo repository.PermissionRepository
}

func NewRoleUseCase(roleRepo repository.RoleRepository, permissionRepo repository.PermissionRepository) RoleUseCase {
	return &roleUseCase{
		roleRepo:       roleRepo,
		permissionRepo: permissionRepo,
	}
}

func (u *roleUseCase) Create(ctx context.Context, role *entity.Role) error {
	return u.roleRepo.Create(role)
}

func (u *roleUseCase) GetByID(ctx context.Context, id uint) (*entity.Role, error) {
	return u.roleRepo.FindByID(id)
}

func (u *roleUseCase) List(ctx context.Context) ([]*entity.Role, error) {
	roles, _, err := u.roleRepo.List(0, 100)
	if err != nil {
		return nil, err
	}

	result := make([]*entity.Role, len(roles))
	for i := range roles {
		result[i] = &roles[i]
	}
	return result, nil
}

func (u *roleUseCase) Update(ctx context.Context, role *entity.Role) error {
	return u.roleRepo.Update(role)
}

func (u *roleUseCase) Delete(ctx context.Context, id uint) error {
	return u.roleRepo.Delete(id)
}

func (u *roleUseCase) AssignPermission(ctx context.Context, roleID uint, permissionID uint) error {
	return u.roleRepo.AssignPermission(roleID, permissionID)
}
