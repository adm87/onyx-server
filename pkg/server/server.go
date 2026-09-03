package server

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

type Server interface {
	Start() error
	Shutdown() error
}

func Run(svr Server, log *zap.Logger) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	errCh := make(chan error, 1)
	go func() {
		errCh <- svr.Start()
	}()

	select {
	case <-signals:
		log.Info("Received shutdown signal, exiting...")
		return svr.Shutdown()

	case err := <-errCh:
		log.Error("server error", zap.Error(err))
		return err
	}
}
