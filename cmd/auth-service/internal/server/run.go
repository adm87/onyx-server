package server

import (
	basicpassword "github.com/adm87/onyx-server/cmd/auth-service/internal/infra/providers/basic_password"
	inmemory "github.com/adm87/onyx-server/cmd/auth-service/internal/infra/repositories/in_memory"
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
	// ===============================================================
	// Compose infrastructure and domain layers

	identityStore := inmemory.NewIdentityStore(cfg, log)
	identityProvider := basicpassword.NewAuthenticator(cfg, log, identityStore)

	// ===============================================================
	// Create gRPC server and register services

	grpcSvr := g.NewServer(cfg.Auth.Svc.Name, &cfg.Auth.Svc.Grpc, log)
	authv1.RegisterAuthServiceServer(grpcSvr.Svr(), v1.NewAuthService(cfg, log, identityProvider))

	// ===============================================================
	// Run the gRPC server

	return g.Run(grpcSvr)
}
