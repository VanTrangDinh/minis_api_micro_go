package email

import (
	"fmt"
	"minisapi/services/auth/configs"
	"net/smtp"
)

type EmailService struct {
	config *configs.EmailConfig
}

func NewEmailService(config *configs.EmailConfig) *EmailService {
	return &EmailService{
		config: config,
	}
}

func (s *EmailService) SendVerificationEmail(to, token string) error {
	subject := "Verify your email"
	body := fmt.Sprintf("Click the link to verify your email: http://localhost:8080/api/v1/auth/verify-email?token=%s", token)
	return s.sendEmail(to, subject, body)
}

func (s *EmailService) SendPasswordResetEmail(to, token string) error {
	subject := "Reset your password"
	body := fmt.Sprintf("Click the link to reset your password: http://localhost:8080/api/v1/auth/reset-password?token=%s", token)
	return s.sendEmail(to, subject, body)
}

func (s *EmailService) sendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body))

	return smtp.SendMail(
		fmt.Sprintf("%s:%d", s.config.Host, s.config.Port),
		auth,
		s.config.From,
		[]string{to},
		msg,
	)
}
