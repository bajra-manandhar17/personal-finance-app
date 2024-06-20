package config

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func NewRedisProvider(ctx context.Context) (*redis.Client, error) {
	redisAddr := os.Getenv("REDIS_ADDR")

	RDB := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return RDB, nil
}
