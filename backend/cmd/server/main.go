package main

import (
	"context"
	"devdeply/internal/adapters/postgres"
	"devdeply/internal/config"
	"devdeply/internal/logger"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)
	logger := logger.New(cfg.App)

	db, err := postgres.New(ctx, cfg.Database, logger)
	if err != nil {
		logger.Error("failed to connect to postgres", "error", err)
		log.Fatal(err)
	}

	defer db.Close()
	logger.Info("application started successfully")
}
