package user

import (
	"context"
	"fmt"
	"net"

	"github.com/adm87/onyx-server/internal/config"
	"github.com/adm87/onyx-server/internal/postgres"
	pgrepository "github.com/adm87/onyx-server/internal/services/user/internal/repository/postgres"
	"github.com/adm87/onyx-server/internal/services/user/internal/server"
	userv1 "github.com/adm87/onyx-server/proto/gen/user/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func Run(cfg *config.Config, log *zap.Logger) error {
	if err := validateConfig(cfg); err != nil {
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Services.User.GRPC.Port))
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

	userRepo := pgrepository.NewUserRepository(db.SQL())

	grpcServer := grpc.NewServer()
	userv1.RegisterUserServiceServer(grpcServer, server.New(cfg, log, userRepo))

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	log.Info("auth service listening", zap.String("port", cfg.Services.User.GRPC.Port))
	return grpcServer.Serve(lis)
}

func validateConfig(cfg *config.Config) error {
	if cfg.Services.User.GRPC.Port == "" {
		return fmt.Errorf("user service grpc port is not set")
	}
	if cfg.Postgres.Schema == "" {
		return fmt.Errorf("postgres schema is not set")
	}
	return nil
}
