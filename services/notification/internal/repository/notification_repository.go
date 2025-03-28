package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"minisapi/services/notification/internal/models"

	"github.com/go-redis/redis/v8"
)

type NotificationRepository interface {
	Create(notification *models.Notification) error
	Update(notification *models.Notification) error
	FindByID(id string) (*models.Notification, error)
	List(page, limit int) ([]models.Notification, int64, error)
}

type notificationRepository struct {
	redis *redis.Client
}

func NewNotificationRepository(redis *redis.Client) NotificationRepository {
	return &notificationRepository{
		redis: redis,
	}
}

// Create saves a new notification to Redis
func (r *notificationRepository) Create(notification *models.Notification) error {
	ctx := context.Background()

	// Convert notification to JSON
	data, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	// Save to Redis
	key := fmt.Sprintf("notification:%s", notification.ID)
	if err := r.redis.Set(ctx, key, data, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("failed to save notification: %w", err)
	}

	// Add to list for pagination
	if err := r.redis.LPush(ctx, "notifications", notification.ID).Err(); err != nil {
		return fmt.Errorf("failed to add notification to list: %w", err)
	}

	return nil
}

// Update updates an existing notification in Redis
func (r *notificationRepository) Update(notification *models.Notification) error {
	ctx := context.Background()

	// Convert notification to JSON
	data, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	// Update in Redis
	key := fmt.Sprintf("notification:%s", notification.ID)
	if err := r.redis.Set(ctx, key, data, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("failed to update notification: %w", err)
	}

	return nil
}

// FindByID retrieves a notification by ID from Redis
func (r *notificationRepository) FindByID(id string) (*models.Notification, error) {
	ctx := context.Background()

	// Get from Redis
	key := fmt.Sprintf("notification:%s", id)
	data, err := r.redis.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get notification: %w", err)
	}

	// Convert JSON to notification
	var notification models.Notification
	if err := json.Unmarshal(data, &notification); err != nil {
		return nil, fmt.Errorf("failed to unmarshal notification: %w", err)
	}

	return &notification, nil
}

// List retrieves a paginated list of notifications from Redis
func (r *notificationRepository) List(page, limit int) ([]models.Notification, int64, error) {
	ctx := context.Background()

	// Get total count
	total, err := r.redis.LLen(ctx, "notifications").Result()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %w", err)
	}

	// Calculate pagination
	start := (page - 1) * limit
	end := start + limit - 1

	// Get IDs from list
	ids, err := r.redis.LRange(ctx, "notifications", int64(start), int64(end)).Result()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get notification IDs: %w", err)
	}

	// Get notifications by IDs
	notifications := make([]models.Notification, 0, len(ids))
	for _, id := range ids {
		notification, err := r.FindByID(id)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to get notification %s: %w", id, err)
		}
		if notification != nil {
			notifications = append(notifications, *notification)
		}
	}

	return notifications, total, nil
}
