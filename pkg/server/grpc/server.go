package grpc

import (
	"context"
	"net"
	"strconv"
	"time"

	"github.com/adm87/onyx-server/pkg/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	cfg    *config.GrpcConfig
	log    *zap.Logger
	server *grpc.Server
	health *health.Server
}

func NewServer(name string, cfg *config.GrpcConfig, log *zap.Logger) *Server {
	health := health.NewServer()
	server := grpc.NewServer()

	healthpb.RegisterHealthServer(server, health)

	health.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	health.SetServingStatus(name, healthpb.HealthCheckResponse_SERVING)

	return &Server{
		cfg:    cfg,
		log:    log,
		server: server,
		health: health,
	}
}

func (s *Server) Svr() *grpc.Server {
	return s.server
}

func (s *Server) Start() error {
	address := net.JoinHostPort("", strconv.Itoa(s.cfg.Port))
	s.log.Info("Starting gRPC server", zap.String("address", address))

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	return s.server.Serve(lis)
}

func (s *Server) Shutdown() error {
	s.log.Info("Shutting down gRPC server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.cfg.ShutdownTimeoutSeconds)*time.Second)
	defer cancel()

	s.health.Shutdown()

	stopped := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		s.log.Info("gRPC server stopped successfully")
	case <-ctx.Done():
		s.log.Error("gRPC server shutdown timed out, forcing stop")
		s.server.Stop()
	}

	return nil
}
