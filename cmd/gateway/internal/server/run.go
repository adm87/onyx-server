package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/adm87/onyx-server/pkg/config"
	g "github.com/adm87/onyx-server/pkg/grpc"
	h "github.com/adm87/onyx-server/pkg/http"
	"github.com/adm87/onyx-server/pkg/logger"
	authv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/auth/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type svcClients struct {
	AuthClient *g.Client
}

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

	svcClients, err := createSvcClients(context.Background(), cfg, log, gw)
	if err != nil {
		return err
	}

	httpSvr, httpErrCh, err := createGateway(cfg, log, gw)
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

		if err := svcClients.AuthClient.Close(); err != nil {
			log.Error("Error closing Auth service connection", zap.Error(err))
			shutdownErrs = append(shutdownErrs, err)
		}

		return errors.Join(shutdownErrs...)

	case err := <-httpErrCh:
		return err
	}
}

func createGateway(cfg *config.Config, log *zap.Logger, gw *runtime.ServeMux) (*h.Server, chan error, error) {
	mux := http.NewServeMux()
	mux.Handle("/", gw)

	httpSvr := h.NewServer(&cfg.Gateway.Http, log, mux)
	httpErrCh := make(chan error, 1)

	go func() {
		httpErrCh <- httpSvr.Start()
	}()

	return httpSvr, httpErrCh, nil
}

func createSvcClients(ctx context.Context, cfg *config.Config, log *zap.Logger, mux *runtime.ServeMux) (*svcClients, error) {
	authClient := g.NewClient(&cfg.Auth.Svc.Grpc, log)
	if err := authClient.Connect(grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		log.Error("Failed to connect to Auth service", zap.Error(err))
		return nil, err
	}

	if err := authv1.RegisterAuthServiceHandler(ctx, mux, authClient.Conn()); err != nil {
		log.Error("Failed to register Auth service handler", zap.Error(err))
		return nil, err
	}

	return &svcClients{
		AuthClient: authClient,
	}, nil
}
