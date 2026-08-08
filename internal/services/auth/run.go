package auth

import (
	"context"
	"fmt"
	"net"

	"github.com/adm87/onyx-server/internal/config"
	"github.com/adm87/onyx-server/internal/postgres"
	pgrepository "github.com/adm87/onyx-server/internal/services/auth/internal/repository/postgres"
	"github.com/adm87/onyx-server/internal/services/auth/internal/server"
	authv1 "github.com/adm87/onyx-server/proto/gen/auth/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func Run(cfg *config.Config, log *zap.Logger) error {
	if err := validateConfig(cfg); err != nil {
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Services.Auth.GRPC.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := postgres.NewDB(&cfg.Postgres, log)
	if err := db.Connect(ctx); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close(ctx)

	identityRepo := pgrepository.NewIdentityRepository(db.SQL())

	grpcServer := grpc.NewServer()
	authv1.RegisterAuthServiceServer(grpcServer, server.New(cfg, log, identityRepo))

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	log.Info("user service listening", zap.String("port", cfg.Services.Auth.GRPC.Port))
	return grpcServer.Serve(lis)
}

func validateConfig(cfg *config.Config) error {
	if cfg.Services.Auth.GRPC.Port == "" {
		return fmt.Errorf("auth service grpc port is not set")
	}
	if cfg.Postgres.Schema == "" {
		return fmt.Errorf("postgres schema is not set")
	}
	return nil
}
