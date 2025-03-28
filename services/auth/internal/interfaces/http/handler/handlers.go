package handler

import (
	"minisapi/services/auth/internal/domain/repository"
	"minisapi/services/auth/internal/domain/service"
	"minisapi/services/auth/internal/domain/usecase"
)

// Handlers contains all handlers
type Handlers struct {
	Auth       *AuthHandler
	Role       *RoleHandler
	Permission *PermissionHandler
}

// NewHandlers creates a new Handlers instance
func NewHandlers(
	userRepo repository.UserRepository,
	roleRepo repository.RoleRepository,
	permissionRepo repository.PermissionRepository,
	tokenRepo repository.TokenRepository,
	sessionRepo repository.SessionRepository,
	authService service.AuthService,
) *Handlers {
	authUseCase := usecase.NewAuthUseCase(userRepo, tokenRepo, sessionRepo, authService)

	return &Handlers{
		Auth:       NewAuthHandler(authUseCase),
		Role:       NewRoleHandler(roleRepo, permissionRepo),
		Permission: NewPermissionHandler(permissionRepo),
	}
}
