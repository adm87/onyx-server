package server

import (
	"github.com/adm87/onyx-server/pkg/config"
	"github.com/adm87/onyx-server/pkg/logger"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	log, err := logger.New(&cfg.Logger)
	if err != nil {
		return err
	}
	log.Info("Starting user service...")
	return nil
}
