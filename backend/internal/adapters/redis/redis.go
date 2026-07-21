package redis

import (
	"context"
	"devdeply/internal/config"
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

func New(ctx context.Context, cfg config.RedisConfig, logger *slog.Logger) (*redis.Client, error) {
	logger.Info("Connecting to Redis...")

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Error("failed to connect to Redis", "error", err)
		return nil, err
	}

	logger.Info("Redis connection established")
	return client, nil
}
