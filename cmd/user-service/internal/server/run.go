package server

import (
	"os"
	"os/signal"
	"syscall"

	v1 "github.com/adm87/onyx-server/cmd/user-service/internal/server/v1"
	"github.com/adm87/onyx-server/pkg/config"
	g "github.com/adm87/onyx-server/pkg/grpc"
	"github.com/adm87/onyx-server/pkg/logger"
	userv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/user/v1"
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
	return run(cfg, log.With(zap.String("prefix", cfg.User.Svc.Name)))
}

func run(cfg *config.Config, log *zap.Logger) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	grpcSvr, grpcErrch, err := createGrpcServer(cfg, log)
	if err != nil {
		return err
	}

	userv1.RegisterUserServiceServer(grpcSvr.Svr(), v1.NewUserSvcRpc(cfg, log))

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
	server := g.NewServer(cfg.User.Svc.Name, &cfg.User.Svc.Grpc, log)
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.Start()
	}()
	return server, errCh, nil
}
