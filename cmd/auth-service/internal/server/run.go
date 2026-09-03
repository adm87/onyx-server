package server

import (
	v1 "github.com/adm87/onyx-server/cmd/auth-service/internal/server/v1"
	"github.com/adm87/onyx-server/cmd/auth-service/internal/store"
	"github.com/adm87/onyx-server/pkg/config"
	"github.com/adm87/onyx-server/pkg/logger"
	authv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/auth/v1"
	"github.com/adm87/onyx-server/pkg/server"
	g "github.com/adm87/onyx-server/pkg/server/grpc"
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
	// ===============================================================
	// Compose infrastructure and domain layers

	identityStore, err := store.NewIdentityStore(cfg, log)
	if err != nil {
		return err
	}
	defer identityStore.Close()

	if err := identityStore.Connect(); err != nil {
		return err
	}

	// ===============================================================
	// Create gRPC server and register services

	grpcSvr := g.NewServer(cfg.Auth.Svc.Name, &cfg.Auth.Svc.Grpc, log)
	authv1.RegisterAuthServiceServer(grpcSvr.Svr(), v1.NewAuthService(cfg, log, identityStore, nil))

	// ===============================================================
	// Run the gRPC server

	return server.Run(grpcSvr, log)
}
