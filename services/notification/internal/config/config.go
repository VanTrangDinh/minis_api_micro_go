package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port string
	Env  string

	// Email
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	SMTPFrom     string

	// SMS
	SMSProvider string
	SMSAPIKey   string
	SMSSecret   string

	// Push Notification
	FirebaseConfig string

	// Redis
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       string

	// RabbitMQ
	RabbitMQHost     string
	RabbitMQPort     string
	RabbitMQUsername string
	RabbitMQPassword string
}

func New() *Config {
	// Load .env file
	_ = godotenv.Load()

	return &Config{
		// Server
		Port: getEnv("APP_PORT", "8082"),
		Env:  getEnv("APP_ENV", "development"),

		// Email
		SMTPHost:     getEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:     getEnv("SMTP_PORT", "587"),
		SMTPUsername: getEnv("SMTP_USERNAME", ""),
		SMTPPassword: getEnv("SMTP_PASSWORD", ""),
		SMTPFrom:     getEnv("SMTP_FROM", ""),

		// SMS
		SMSProvider: getEnv("SMS_PROVIDER", "twilio"),
		SMSAPIKey:   getEnv("SMS_API_KEY", ""),
		SMSSecret:   getEnv("SMS_SECRET", ""),

		// Push Notification
		FirebaseConfig: getEnv("FIREBASE_CONFIG", ""),

		// Redis
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnv("REDIS_DB", "0"),

		// RabbitMQ
		RabbitMQHost:     getEnv("RABBITMQ_HOST", "localhost"),
		RabbitMQPort:     getEnv("RABBITMQ_PORT", "5672"),
		RabbitMQUsername: getEnv("RABBITMQ_USERNAME", "guest"),
		RabbitMQPassword: getEnv("RABBITMQ_PASSWORD", "guest"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
