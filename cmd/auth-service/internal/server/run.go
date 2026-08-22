package server

import (
	"os"
	"os/signal"
	"syscall"

	v1 "github.com/adm87/onyx-server/cmd/auth-service/internal/server/v1"
	"github.com/adm87/onyx-server/pkg/config"
	g "github.com/adm87/onyx-server/pkg/grpc"
	"github.com/adm87/onyx-server/pkg/logger"
	authv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/auth/v1"
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
	return run(cfg, log.With(zap.String("prefix", cfg.Auth.Svc.Name)))
}

func run(cfg *config.Config, log *zap.Logger) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	grpcSvr, grpcErrch, err := createGrpcServer(cfg, log)
	if err != nil {
		return err
	}

	select {
	case <-signals:
		log.Info("Received shutdown signal, exiting...")

		if err := grpcSvr.Shutdown(); err != nil {
			log.Error("Error shutting down gRPC server", zap.Error(err))
			return err
		}
		return nil

	case err := <-grpcErrch:
		return err
	}
}

func createGrpcServer(cfg *config.Config, log *zap.Logger) (*g.Server, chan error, error) {
	server := g.NewServer(cfg.Auth.Svc.Name, &cfg.Auth.Svc.Grpc, log)

	authv1.RegisterAuthServiceServer(server.Svr(), v1.NewAuthSvcRpc(cfg, log))

	errCh := make(chan error, 1)

	go func() {
		errCh <- server.Start()
	}()

	return server, errCh, nil
}
