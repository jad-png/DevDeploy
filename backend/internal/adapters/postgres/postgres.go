package postgres

import (
	"context"
	"devdeply/internal/config"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func New(ctx context.Context, cfg config.DatabaseConfig, logger *slog.Logger) (*pgxpool.Pool, error) {
	logger.Info("Connecting to Postgres...")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Error("failed to parse PostgreSQL configuration", "error", err)
		return nil, err
	}

	poolConfig.MaxConns = cfg.Pool.MaxConns
	poolConfig.MinConns = cfg.Pool.MinConns
	poolConfig.MaxConnIdleTime = cfg.Pool.MaxConnIdleTime
	poolConfig.MaxConnLifetime = cfg.Pool.MaxConnLifetime

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		logger.Error("failed to create PostgreSQL connection pool", "error", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		logger.Error("failed to connect to PostgreSQL", "error", err)
		return nil, err
	}

	logger.Info("PostgreSQL connection established")

	return pool, nil
}
