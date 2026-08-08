package main

import (
	"log"

	"github.com/adm87/onyx-server/internal/config"
	"github.com/adm87/onyx-server/internal/logging"
	"github.com/adm87/onyx-server/internal/services/auth"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	lgr, err := logging.New(&cfg.Logger)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer lgr.Sync()

	if err := auth.Run(cfg, lgr); err != nil {
		lgr.Fatal("exiting with error", zap.Error(err))
	}
}
