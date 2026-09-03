package server

import (
	"context"
	"net/http"

	"github.com/adm87/onyx-server/cmd/gateway/internal/openapi"
	"github.com/adm87/onyx-server/pkg/config"
	"github.com/adm87/onyx-server/pkg/logger"
	"github.com/adm87/onyx-server/pkg/server"
	h "github.com/adm87/onyx-server/pkg/server/http"
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
	gw := runtime.NewServeMux()

	clients, err := newSvcClients(context.Background(), cfg, log, gw)
	if err != nil {
		return err
	}
	defer clients.Close()

	mux := http.NewServeMux()
	mux.Handle("/", gw)
	mux.HandleFunc("/healthz", healthzHandler(clients, log))

	if cfg.Gateway.EnableSwaggerUI {
		mux.Handle("/docs", openapi.RegisterSwaggerUI(mux, log))
	}

	httpServer := h.NewServer(&cfg.Gateway.Http, log, mux)
	return server.Run(httpServer, log)
}
