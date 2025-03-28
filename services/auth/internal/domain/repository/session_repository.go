package repository

import (
	"minisapi/services/auth/internal/domain/entity"
)

type SessionRepository interface {
	Create(session *entity.Session) error
	FindByID(id uint) (*entity.Session, error)
	FindByUserID(userID uint) ([]entity.Session, error)
	FindByTokenID(tokenID uint) (*entity.Session, error)
	Update(session *entity.Session) error
	Delete(id uint) error
	Deactivate(id uint) error
	DeactivateAll(userID uint) error
	CleanupExpired() error
	CountActiveSessions(userID uint) (int64, error)
}
