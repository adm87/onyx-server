package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/adm87/onyx-server/pkg/config"
	h "github.com/adm87/onyx-server/pkg/http"
	"github.com/adm87/onyx-server/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
	return run(cfg, log.With(zap.String("prefix", cfg.Gateway.Name)))
}

func run(cfg *config.Config, log *zap.Logger) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	gw := runtime.NewServeMux()

	clients, err := createSvcClients(context.Background(), cfg, log, gw)
	if err != nil {
		return err
	}

	httpSvr, httpErrCh, err := createGateway(cfg, log, gw, clients)
	if err != nil {
		return err
	}

	select {
	case <-signals:
		log.Info("Received shutdown signal, exiting...")

		var shutdownErrs []error
		if err := httpSvr.Shutdown(); err != nil {
			log.Error("Error shutting down HTTP server", zap.Error(err))
			shutdownErrs = append(shutdownErrs, err)
		}

		if err := clients.Close(); err != nil {
			log.Error("Error closing Auth service connection", zap.Error(err))
			shutdownErrs = append(shutdownErrs, err)
		}

		return errors.Join(shutdownErrs...)

	case err := <-httpErrCh:
		return err
	}
}

func createGateway(cfg *config.Config, log *zap.Logger, gw *runtime.ServeMux, clients *svcClients) (*h.Server, chan error, error) {
	mux := http.NewServeMux()

	mux.Handle("/", gw)
	mux.HandleFunc("/healthz", healthzHandler(clients, log))

	httpSvr := h.NewServer(&cfg.Gateway.Http, log, mux)
	httpErrCh := make(chan error, 1)

	go func() {
		httpErrCh <- httpSvr.Start()
	}()

	return httpSvr, httpErrCh, nil
}
