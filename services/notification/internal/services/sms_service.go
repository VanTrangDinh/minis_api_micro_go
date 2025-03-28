package services

import (
	"fmt"
	"time"

	"minisapi/services/notification/internal/config"
	"minisapi/services/notification/internal/models"
	"minisapi/services/notification/internal/repository"
)

type SMSService struct {
	config *config.Config
	repo   repository.NotificationRepository
}

func NewSMSService(cfg *config.Config) *SMSService {
	return &SMSService{
		config: cfg,
	}
}

// SendSMS sends an SMS notification
func (s *SMSService) SendSMS(req *models.SMSRequest) (*models.Notification, error) {
	// Create notification record
	notification := &models.Notification{
		Type:      models.SMSNotification,
		Status:    models.StatusPending,
		Recipient: req.To,
		Content:   req.Content,
		Metadata:  req.Metadata,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to repository
	if err := s.repo.Create(notification); err != nil {
		return nil, fmt.Errorf("failed to create notification: %w", err)
	}

	// Send SMS based on provider
	if err := s.sendSMS(req); err != nil {
		// Update notification status
		notification.Status = models.StatusFailed
		errorMsg := err.Error()
		notification.Error = &errorMsg
		notification.UpdatedAt = time.Now()
		s.repo.Update(notification)
		return nil, fmt.Errorf("failed to send SMS: %w", err)
	}

	// Update notification status
	notification.Status = models.StatusSent
	sentAt := time.Now()
	notification.SentAt = &sentAt
	notification.UpdatedAt = time.Now()
	s.repo.Update(notification)

	return notification, nil
}

// sendSMS sends the actual SMS using the configured provider
func (s *SMSService) sendSMS(req *models.SMSRequest) error {
	switch s.config.SMSProvider {
	case "twilio":
		return s.sendTwilioSMS(req)
	case "nexmo":
		return s.sendNexmoSMS(req)
	default:
		return fmt.Errorf("unsupported SMS provider: %s", s.config.SMSProvider)
	}
}

// sendTwilioSMS sends SMS using Twilio
func (s *SMSService) sendTwilioSMS(req *models.SMSRequest) error {
	// TODO: Implement Twilio SMS sending
	// This is a placeholder for actual implementation
	return fmt.Errorf("twilio SMS sending not implemented")
}

// sendNexmoSMS sends SMS using Nexmo
func (s *SMSService) sendNexmoSMS(req *models.SMSRequest) error {
	// TODO: Implement Nexmo SMS sending
	// This is a placeholder for actual implementation
	return fmt.Errorf("nexmo SMS sending not implemented")
}
