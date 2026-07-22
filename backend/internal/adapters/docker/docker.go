package docker

import (
	"context"
	"devdeply/internal/config"
	"log/slog"

	"github.com/moby/moby/client"
)

func New(ctx context.Context, cfg config.DockerConfig, logger *slog.Logger) (*client.Client, error) {
	logger.Info("Connecting to Docker Engine...")

	cli, err := client.New(
		client.WithHost(cfg.Host),
	)
	if err != nil {
		logger.Error("failed to create Docker client", "error", err)
		return nil, err
	}

	if _, err := cli.ServerVersion(ctx, client.ServerVersionOptions{}); err != nil {
		logger.Error("failed to connect to Docker Engine", "error", err)
		_ = cli.Close()
		return nil, err
	}

	logger.Info("Docker Engine connection established")

	return cli, nil
}
