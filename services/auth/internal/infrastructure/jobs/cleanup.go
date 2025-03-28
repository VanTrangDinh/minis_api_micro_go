package jobs

import (
	"context"
	"time"

	"minisapi/services/auth/internal/domain/repository"
)

type CleanupJob struct {
	tokenRepo   repository.TokenRepository
	sessionRepo repository.SessionRepository
}

func NewCleanupJob(tokenRepo repository.TokenRepository, sessionRepo repository.SessionRepository) *CleanupJob {
	return &CleanupJob{
		tokenRepo:   tokenRepo,
		sessionRepo: sessionRepo,
	}
}

func (j *CleanupJob) Start(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := j.cleanup(ctx); err != nil {
				// TODO: Log error
			}
		}
	}
}

func (j *CleanupJob) cleanup(ctx context.Context) error {
	// Cleanup expired tokens
	if err := j.tokenRepo.CleanupExpired(); err != nil {
		return err
	}

	// Cleanup expired sessions
	if err := j.sessionRepo.CleanupExpired(); err != nil {
		return err
	}

	return nil
}
