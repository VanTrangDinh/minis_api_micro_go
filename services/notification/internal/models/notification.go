package models

import (
	"time"
)

// NotificationType represents the type of notification
type NotificationType string

const (
	EmailNotification NotificationType = "email"
	SMSNotification   NotificationType = "sms"
	PushNotification  NotificationType = "push"
)

// NotificationStatus represents the status of a notification
type NotificationStatus string

const (
	StatusPending   NotificationStatus = "pending"
	StatusSent      NotificationStatus = "sent"
	StatusFailed    NotificationStatus = "failed"
	StatusDelivered NotificationStatus = "delivered"
)

// Notification represents a notification in the system
type Notification struct {
	ID        string             `json:"id"`
	Type      NotificationType   `json:"type"`
	Status    NotificationStatus `json:"status"`
	Recipient string             `json:"recipient"`
	Subject   string             `json:"subject,omitempty"`
	Content   string             `json:"content"`
	Metadata  map[string]string  `json:"metadata,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	SentAt    *time.Time         `json:"sent_at,omitempty"`
	Error     *string            `json:"error,omitempty"`
}

// EmailRequest represents an email notification request
type EmailRequest struct {
	To          []string          `json:"to" binding:"required"`
	Cc          []string          `json:"cc,omitempty"`
	Bcc         []string          `json:"bcc,omitempty"`
	Subject     string            `json:"subject" binding:"required"`
	Content     string            `json:"content" binding:"required"`
	HTML        bool              `json:"html,omitempty"`
	Attachments []Attachment      `json:"attachments,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// SMSRequest represents an SMS notification request
type SMSRequest struct {
	To       string            `json:"to" binding:"required"`
	Content  string            `json:"content" binding:"required"`
	From     string            `json:"from,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// PushRequest represents a push notification request
type PushRequest struct {
	DeviceToken string            `json:"device_token" binding:"required"`
	Title       string            `json:"title" binding:"required"`
	Body        string            `json:"body" binding:"required"`
	Data        map[string]string `json:"data,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// Attachment represents an email attachment
type Attachment struct {
	Filename    string `json:"filename" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

// NotificationList represents a paginated list of notifications
type NotificationList struct {
	Items      []Notification `json:"items"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"total_pages"`
}
