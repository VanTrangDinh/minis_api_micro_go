package repository

import (
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/domain/repository"
	"minisapi/services/auth/internal/pkg/errors"
	"time"

	"gorm.io/gorm"
)

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) repository.TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) Create(token *entity.Token) error {
	if err := r.db.Create(token).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *tokenRepository) FindByToken(token string) (*entity.Token, error) {
	var t entity.Token
	if err := r.db.Where("token = ?", token).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrTokenNotFound
		}
		return nil, errors.ErrDatabase
	}
	return &t, nil
}

func (r *tokenRepository) FindByUserID(userID uint) ([]entity.Token, error) {
	var tokens []entity.Token
	if err := r.db.Where("user_id = ?", userID).Find(&tokens).Error; err != nil {
		return nil, errors.ErrDatabase
	}
	return tokens, nil
}

func (r *tokenRepository) Update(token *entity.Token) error {
	if err := r.db.Save(token).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *tokenRepository) Delete(id uint) error {
	if err := r.db.Delete(&entity.Token{}, id).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *tokenRepository) Revoke(id uint) error {
	if err := r.db.Model(&entity.Token{}).Where("id = ?", id).Update("revoked", true).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *tokenRepository) RevokeAll(userID uint) error {
	if err := r.db.Model(&entity.Token{}).Where("user_id = ?", userID).Update("revoked", true).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *tokenRepository) CleanupExpired() error {
	if err := r.db.Where("expires_at < ?", time.Now()).Delete(&entity.Token{}).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *tokenRepository) FindByID(id uint) (*entity.Token, error) {
	var t entity.Token
	if err := r.db.First(&t, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrTokenNotFound
		}
		return nil, errors.ErrDatabase
	}
	return &t, nil
}

func (r *tokenRepository) RevokeByType(userID uint, tokenType entity.TokenType) error {
	if err := r.db.Model(&entity.Token{}).Where("user_id = ? AND type = ?", userID, tokenType).Update("revoked", true).Error; err != nil {
		return errors.ErrDatabase
	}
	return nil
}
