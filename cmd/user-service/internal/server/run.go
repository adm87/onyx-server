package server

import (
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
	// ===============================================================
	// Create gRPC server and register services

	grpcSvr := g.NewServer(cfg.User.Svc.Name, &cfg.User.Svc.Grpc, log)
	userv1.RegisterUserServiceServer(grpcSvr.Svr(), v1.NewUserService(cfg, log))

	// ===============================================================
	// Run the gRPC server

	return g.Run(grpcSvr)
}
