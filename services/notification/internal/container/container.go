package container

import (
	"minisapi/services/notification/internal/config"
	"minisapi/services/notification/internal/repository"
	"minisapi/services/notification/internal/services"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
)

// Container holds all dependencies
type Container struct {
	config *config.Config

	// Redis client
	redis *redis.Client

	// RabbitMQ connection
	rabbitMQ *amqp.Connection

	// Repositories
	notificationRepo repository.NotificationRepository

	// Services
	emailService *services.EmailService
	smsService   *services.SMSService
	pushService  *services.PushService
	queueService *services.QueueService
}

// NewContainer creates a new container with all dependencies
func NewContainer(cfg *config.Config) *Container {
	container := &Container{
		config: cfg,
	}

	// Initialize Redis
	container.redis = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	// Initialize RabbitMQ
	amqpURL := "amqp://" + cfg.RabbitMQUsername + ":" + cfg.RabbitMQPassword + "@" + cfg.RabbitMQHost + ":" + cfg.RabbitMQPort + "/"
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		panic(err)
	}
	container.rabbitMQ = conn

	// Initialize repositories
	container.notificationRepo = repository.NewNotificationRepository(container.redis)

	// Initialize services
	container.emailService = services.NewEmailService(cfg)
	container.smsService = services.NewSMSService(cfg)
	container.pushService = services.NewPushService(cfg)
	container.queueService = services.NewQueueService(container.rabbitMQ)

	return container
}

// GetConfig returns the config
func (c *Container) GetConfig() *config.Config {
	return c.config
}

// GetRedis returns the Redis client
func (c *Container) GetRedis() *redis.Client {
	return c.redis
}

// GetRabbitMQ returns the RabbitMQ connection
func (c *Container) GetRabbitMQ() *amqp.Connection {
	return c.rabbitMQ
}

// GetNotificationRepo returns the notification repository
func (c *Container) GetNotificationRepo() repository.NotificationRepository {
	return c.notificationRepo
}

// GetEmailService returns the email service
func (c *Container) GetEmailService() *services.EmailService {
	return c.emailService
}

// GetSMSService returns the SMS service
func (c *Container) GetSMSService() *services.SMSService {
	return c.smsService
}

// GetPushService returns the push notification service
func (c *Container) GetPushService() *services.PushService {
	return c.pushService
}

// GetQueueService returns the queue service
func (c *Container) GetQueueService() *services.QueueService {
	return c.queueService
}
