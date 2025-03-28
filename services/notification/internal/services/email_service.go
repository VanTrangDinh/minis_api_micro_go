package services

import (
	"fmt"
	"net/smtp"
	"time"

	"minisapi/services/notification/internal/config"
	"minisapi/services/notification/internal/models"
	"minisapi/services/notification/internal/repository"
)

type EmailService struct {
	config *config.Config
	repo   repository.NotificationRepository
}

func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{
		config: cfg,
	}
}

// SendEmail sends an email notification
func (s *EmailService) SendEmail(req *models.EmailRequest) (*models.Notification, error) {
	// Create notification record
	notification := &models.Notification{
		Type:      models.EmailNotification,
		Status:    models.StatusPending,
		Recipient: req.To[0], // Primary recipient
		Subject:   req.Subject,
		Content:   req.Content,
		Metadata:  req.Metadata,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to repository
	if err := s.repo.Create(notification); err != nil {
		return nil, fmt.Errorf("failed to create notification: %w", err)
	}

	// Send email
	if err := s.sendEmail(req); err != nil {
		// Update notification status
		notification.Status = models.StatusFailed
		errorMsg := err.Error()
		notification.Error = &errorMsg
		notification.UpdatedAt = time.Now()
		s.repo.Update(notification)
		return nil, fmt.Errorf("failed to send email: %w", err)
	}

	// Update notification status
	notification.Status = models.StatusSent
	sentAt := time.Now()
	notification.SentAt = &sentAt
	notification.UpdatedAt = time.Now()
	s.repo.Update(notification)

	return notification, nil
}

// sendEmail sends the actual email using SMTP
func (s *EmailService) sendEmail(req *models.EmailRequest) error {
	// Set up authentication information
	auth := smtp.PlainAuth("",
		s.config.SMTPUsername,
		s.config.SMTPPassword,
		s.config.SMTPHost,
	)

	// Prepare email headers
	headers := make(map[string]string)
	headers["From"] = s.config.SMTPFrom
	headers["To"] = req.To[0]
	headers["Subject"] = req.Subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""

	// Build message
	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + req.Content

	// Send email
	return smtp.SendMail(
		fmt.Sprintf("%s:%s", s.config.SMTPHost, s.config.SMTPPort),
		auth,
		s.config.SMTPFrom,
		req.To,
		[]byte(message),
	)
}
