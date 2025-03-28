package redis

import (
	"context"
	"fmt"
	"minisapi/services/auth/internal/configs"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg configs.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	})

	// Test connection
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return client, nil
}
