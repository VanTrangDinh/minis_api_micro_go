package repository

import (
	"context"
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/domain/repository"
	"minisapi/services/auth/internal/pkg/errors"
	"time"

	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) repository.SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) Create(session *entity.Session) error {
	result := r.db.Create(session)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *sessionRepository) FindByToken(token string) (*entity.Session, error) {
	var session entity.Session
	result := r.db.Where("token = ?", token).First(&session)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrSessionNotFound
		}
		return nil, errors.ErrDatabase
	}
	return &session, nil
}

func (r *sessionRepository) DeleteByToken(ctx context.Context, token string) error {
	result := r.db.WithContext(ctx).Where("token = ?", token).Delete(&entity.Session{})
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *sessionRepository) DeleteExpired(ctx context.Context) error {
	result := r.db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&entity.Session{})
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *sessionRepository) CountByUserID(ctx context.Context, userID uint) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&entity.Session{}).Where("user_id = ?", userID).Count(&count)
	if result.Error != nil {
		return 0, errors.ErrDatabase
	}
	return count, nil
}

func (r *sessionRepository) CleanupExpired() error {
	return r.DeleteExpired(context.Background())
}

func (r *sessionRepository) CountActiveSessions(userID uint) (int64, error) {
	var count int64
	result := r.db.Model(&entity.Session{}).Where("user_id = ? AND expires_at > ?", userID, time.Now()).Count(&count)
	if result.Error != nil {
		return 0, errors.ErrDatabase
	}
	return count, nil
}

func (r *sessionRepository) Deactivate(userID uint) error {
	result := r.db.Model(&entity.Session{}).Where("user_id = ?", userID).Update("active", false)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *sessionRepository) DeactivateAll(userID uint) error {
	result := r.db.Model(&entity.Session{}).Where("user_id = ?", userID).Update("active", false)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *sessionRepository) Delete(id uint) error {
	result := r.db.Delete(&entity.Session{}, id)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}

func (r *sessionRepository) FindByID(id uint) (*entity.Session, error) {
	var session entity.Session
	result := r.db.First(&session, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrSessionNotFound
		}
		return nil, errors.ErrDatabase
	}
	return &session, nil
}

func (r *sessionRepository) FindByUserID(userID uint) ([]entity.Session, error) {
	var sessions []entity.Session
	result := r.db.Where("user_id = ?", userID).Find(&sessions)
	if result.Error != nil {
		return nil, errors.ErrDatabase
	}
	return sessions, nil
}

func (r *sessionRepository) FindByTokenID(tokenID uint) (*entity.Session, error) {
	var session entity.Session
	result := r.db.Where("token_id = ?", tokenID).First(&session)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrSessionNotFound
		}
		return nil, errors.ErrDatabase
	}
	return &session, nil
}

func (r *sessionRepository) Update(session *entity.Session) error {
	result := r.db.Save(session)
	if result.Error != nil {
		return errors.ErrDatabase
	}
	return nil
}
