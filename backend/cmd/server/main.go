package main

import (
	"devdeply/internal/config"
	"devdeply/internal/logger"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logg := logger.New(cfg.App)

	logg.Error("test app")
}
