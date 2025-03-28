package errors

import "errors"

var (
	// Authentication errors
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAccountLocked      = errors.New("account is locked")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token has expired")
	ErrTokenRevoked       = errors.New("token has been revoked")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrPasswordHash       = errors.New("failed to hash password")
	ErrInvalidTwoFactor   = errors.New("invalid two factor code")

	// User errors
	ErrUsernameExists    = errors.New("username already exists")
	ErrEmailExists       = errors.New("email already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrUserInactive      = errors.New("user is inactive")
	ErrUserDeleted       = errors.New("user has been deleted")
	ErrEmailNotVerified  = errors.New("email not verified")
	ErrTwoFactorRequired = errors.New("two factor authentication required")

	// Role errors
	ErrRoleNotFound = errors.New("role not found")
	ErrRoleExists   = errors.New("role already exists")
	ErrInvalidRole  = errors.New("invalid role")
	ErrRoleInUse    = errors.New("role is in use")

	// Permission errors
	ErrPermissionNotFound = errors.New("permission not found")
	ErrPermissionExists   = errors.New("permission already exists")
	ErrInvalidPermission  = errors.New("invalid permission")
	ErrPermissionDenied   = errors.New("permission denied")

	// Session errors
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionExpired  = errors.New("session has expired")
	ErrSessionRevoked  = errors.New("session has been revoked")
	ErrTooManySessions = errors.New("too many active sessions")

	// Validation errors
	ErrValidation   = errors.New("validation error")
	ErrInvalidInput = errors.New("invalid input")

	// Database errors
	ErrDatabase       = errors.New("database error")
	ErrDuplicateEntry = errors.New("duplicate entry")

	// Server errors
	ErrInternalServer     = errors.New("internal server error")
	ErrServiceUnavailable = errors.New("service unavailable")

	// User errors
	ErrUserExists        = errors.New("user already exists")
	ErrUserLocked        = errors.New("user account is locked")

	// Token errors
	ErrTokenNotFound     = errors.New("token not found")
	ErrTokenInvalid      = errors.New("invalid token")
	ErrTokenTypeInvalid  = errors.New("invalid token type")

	// Session errors
	ErrSessionInvalid    = errors.New("invalid session")

	// Rate limiting errors
	ErrRateLimitExceeded = errors.New("rate limit exceeded")

	// Authentication errors
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
)
