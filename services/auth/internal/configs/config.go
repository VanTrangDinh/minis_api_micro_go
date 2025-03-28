package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
	Log      LogConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	Timezone string
}

type RedisConfig struct {
	Host string
	Port string
}

type JWTConfig struct {
	Secret     string
	Expiration string
}

type LogConfig struct {
	Level string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnvOrDefault("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvOrDefault("DB_PORT", "5432"),
			User:     getEnvOrDefault("DB_USER", "postgres"),
			Password: getEnvOrDefault("DB_PASSWORD", "postgres"),
			DBName:   getEnvOrDefault("DB_NAME", "minisapi_auth"),
			SSLMode:  getEnvOrDefault("DB_SSLMODE", "disable"),
			Timezone: getEnvOrDefault("DB_TIMEZONE", "Asia/Ho_Chi_Minh"),
		},
		Redis: RedisConfig{
			Host: getEnvOrDefault("REDIS_HOST", "localhost"),
			Port: getEnvOrDefault("REDIS_PORT", "6379"),
		},
		JWT: JWTConfig{
			Secret:     getEnvOrDefault("JWT_SECRET", "your-secret-key"),
			Expiration: getEnvOrDefault("JWT_EXPIRATION", "24h"),
		},
		Log: LogConfig{
			Level: getEnvOrDefault("LOG_LEVEL", "info"),
		},
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
