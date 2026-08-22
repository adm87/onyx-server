package server

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/adm87/onyx-server/pkg/config"
	h "github.com/adm87/onyx-server/pkg/http"
	"github.com/adm87/onyx-server/pkg/logger"
	"go.uber.org/zap"
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
	return run(cfg, log.With(zap.String("prefix", "gateway")))
}

func run(cfg *config.Config, log *zap.Logger) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	httpSvr, httpErrChan, err := runHttpServer(cfg, log)
	if err != nil {
		return err
	}

	select {
	case <-signals:
		log.Info("Received shutdown signal, exiting...")

		if err := httpSvr.Shutdown(); err != nil {
			log.Error("Error shutting down HTTP server", zap.Error(err))
			return err
		}
		return nil

	case err := <-httpErrChan:
		return err
	}
}

func runHttpServer(cfg *config.Config, log *zap.Logger) (*h.Server, chan error, error) {
	mux := http.NewServeMux()
	svr := h.NewServer(&cfg.Gateway, log, mux)

	errChan := make(chan error, 1)
	go func() {
		errChan <- svr.Start()
	}()
	return svr, errChan, nil
}
