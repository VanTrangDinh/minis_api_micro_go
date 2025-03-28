package services

import (
	"fmt"
	"time"

	"minisapi/services/notification/internal/config"
	"minisapi/services/notification/internal/models"
	"minisapi/services/notification/internal/repository"
)

type PushService struct {
	config *config.Config
	repo   repository.NotificationRepository
}

func NewPushService(cfg *config.Config) *PushService {
	return &PushService{
		config: cfg,
	}
}

// SendPush sends a push notification
func (s *PushService) SendPush(req *models.PushRequest) (*models.Notification, error) {
	// Create notification record
	notification := &models.Notification{
		Type:      models.PushNotification,
		Status:    models.StatusPending,
		Recipient: req.DeviceToken,
		Subject:   req.Title,
		Content:   req.Body,
		Metadata:  req.Metadata,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to repository
	if err := s.repo.Create(notification); err != nil {
		return nil, fmt.Errorf("failed to create notification: %w", err)
	}

	// Send push notification
	if err := s.sendPush(req); err != nil {
		// Update notification status
		notification.Status = models.StatusFailed
		errorMsg := err.Error()
		notification.Error = &errorMsg
		notification.UpdatedAt = time.Now()
		s.repo.Update(notification)
		return nil, fmt.Errorf("failed to send push notification: %w", err)
	}

	// Update notification status
	notification.Status = models.StatusSent
	sentAt := time.Now()
	notification.SentAt = &sentAt
	notification.UpdatedAt = time.Now()
	s.repo.Update(notification)

	return notification, nil
}

// sendPush sends the actual push notification using Firebase
func (s *PushService) sendPush(req *models.PushRequest) error {
	// TODO: Implement Firebase push notification
	// This is a placeholder for actual implementation
	return fmt.Errorf("push notification not implemented")
}
