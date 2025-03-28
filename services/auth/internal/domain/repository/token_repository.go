package repository

import (
	"minisapi/services/auth/internal/domain/entity"
)

type TokenRepository interface {
	Create(token *entity.Token) error
	FindByID(id uint) (*entity.Token, error)
	FindByToken(token string) (*entity.Token, error)
	FindByUserID(userID uint) ([]entity.Token, error)
	Update(token *entity.Token) error
	Delete(id uint) error
	Revoke(id uint) error
	RevokeAll(userID uint) error
	RevokeByType(userID uint, tokenType entity.TokenType) error
	CleanupExpired() error
}
