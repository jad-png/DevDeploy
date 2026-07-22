package main

import (
	"context"
	"devdeply/internal/adapters/docker"
	"devdeply/internal/adapters/postgres"
	"devdeply/internal/adapters/redis"
	"devdeply/internal/config"
	"devdeply/internal/logger"
	"io"
	"log"
	"log/slog"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logg := logger.New(cfg.App)

	db, err := postgres.New(ctx, cfg.Database, logg)
	if err != nil {
		logg.Error("failed to connect to postgres", "error", err)
		log.Fatal(err)
	}

	redisClient, err := redis.New(ctx, cfg.Redis, logg)
	if err != nil {
		logg.Error("failed to connect to redis", "error", err)
		log.Fatal(err)
	}

	dockerClient, err := docker.New(ctx, cfg.Docker, logg)
	if err != nil {
		logg.Error("failed to connect to docker", "error", err)
		log.Fatal(err)
	}

	defer db.Close()

	defer closeResources(
		logg,
		redisClient,
		dockerClient,
	)

	logg.Info("application started successfully")
}

func closeResources(logger *slog.Logger, closers ...io.Closer) {
	for _, c := range closers {
		if err := c.Close(); err != nil {
			logger.Error("failed to close resource", "error", err)
		}
	}
}
