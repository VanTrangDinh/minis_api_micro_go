package usecase

import (
	"context"
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/domain/repository"
	"minisapi/services/auth/internal/domain/service"
)

type AuthUseCase interface {
	Register(ctx context.Context, user *entity.User) error
	Login(ctx context.Context, email, password string) (*entity.User, string, string, error)
	Logout(ctx context.Context, userID uint, refreshToken string) error
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
	ResetPassword(ctx context.Context, email string) error
	VerifyEmail(ctx context.Context, token string) error
	ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error
	EnableTwoFactor(ctx context.Context, userID uint) (string, error)
	DisableTwoFactor(ctx context.Context, userID uint) error
	VerifyTwoFactor(ctx context.Context, userID uint, code string) error
	ForgotPassword(ctx context.Context, email string) error
}

type authUseCase struct {
	userRepo    repository.UserRepository
	tokenRepo   repository.TokenRepository
	sessionRepo repository.SessionRepository
	authService service.AuthService
}

func (uc *authUseCase) Register(ctx context.Context, user *entity.User) error {
	return uc.userRepo.Create(user)
}

func (uc *authUseCase) Login(ctx context.Context, email, password string) (*entity.User, string, string, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", "", err
	}

	if err := uc.authService.ValidatePassword(user, password); err != nil {
		return nil, "", "", err
	}

	accessToken, refreshToken, err := uc.authService.GenerateTokens(user)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

func (uc *authUseCase) Logout(ctx context.Context, userID uint, refreshToken string) error {
	if err := uc.tokenRepo.RevokeAll(userID); err != nil {
		return err
	}

	token, err := uc.tokenRepo.FindByToken(refreshToken)
	if err != nil {
		return err
	}

	session, err := uc.sessionRepo.FindByTokenID(token.ID)
	if err != nil {
		return err
	}

	return uc.sessionRepo.Deactivate(session.ID)
}

func (uc *authUseCase) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	token, err := uc.tokenRepo.FindByToken(refreshToken)
	if err != nil {
		return "", err
	}

	user, err := uc.userRepo.FindByID(token.UserID)
	if err != nil {
		return "", err
	}

	return uc.authService.GenerateAccessToken(user)
}

func (uc *authUseCase) ResetPassword(ctx context.Context, email string) error {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return err
	}

	return uc.authService.SendPasswordResetEmail(user)
}

func (uc *authUseCase) VerifyEmail(ctx context.Context, token string) error {
	return uc.authService.VerifyEmail(token)
}

func (uc *authUseCase) ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error {
	user, err := uc.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	if err := uc.authService.ValidatePassword(user, oldPassword); err != nil {
		return err
	}

	return uc.userRepo.UpdatePassword(userID, newPassword)
}

func (uc *authUseCase) EnableTwoFactor(ctx context.Context, userID uint) (string, error) {
	secret, err := uc.authService.GenerateTwoFactorSecret()
	if err != nil {
		return "", err
	}

	err = uc.userRepo.UpdateTwoFactor(ctx, userID, true, secret)
	if err != nil {
		return "", err
	}

	return secret, nil
}

func (uc *authUseCase) DisableTwoFactor(ctx context.Context, userID uint) error {
	return uc.userRepo.UpdateTwoFactor(ctx, userID, false, "")
}

func (uc *authUseCase) VerifyTwoFactor(ctx context.Context, userID uint, code string) error {
	user, err := uc.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	return uc.authService.ValidateTwoFactorCode(user, code)
	}

func (uc *authUseCase) ForgotPassword(ctx context.Context, email string) error {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return err
	}

	return uc.authService.SendPasswordResetEmail(user)
}

func NewAuthUseCase(
	userRepo repository.UserRepository,
	tokenRepo repository.TokenRepository,
	sessionRepo repository.SessionRepository,
	authService service.AuthService,
) AuthUseCase {
	return &authUseCase{
		userRepo:    userRepo,
		tokenRepo:   tokenRepo,
		sessionRepo: sessionRepo,
		authService: authService,
	}
}
