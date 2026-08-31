package http

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func Run(server *Server) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.Start()
	}()

	select {
	case <-signals:
		server.log.Info("Received shutdown signal, exiting...")
		return server.Shutdown()

	case err := <-errCh:
		server.log.Error("HTTP server exited with error", zap.Error(err))
		return err
	}
}
