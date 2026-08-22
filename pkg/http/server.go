package http

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/adm87/onyx-server/pkg/config"
	"go.uber.org/zap"
)

type Server struct {
	cfg    *config.HttpConfig
	log    *zap.Logger
	server *http.Server
}

func NewServer(cfg *config.HttpConfig, log *zap.Logger, handler http.Handler) *Server {
	addr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	return &Server{
		cfg: cfg,
		log: log,
		server: &http.Server{
			Addr:         addr,
			Handler:      handler,
			IdleTimeout:  time.Duration(cfg.IdleTimeoutSeconds) * time.Second,
			ReadTimeout:  time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
			WriteTimeout: time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
		},
	}
}

func (s *Server) Start() error {
	s.log.Info("Starting HTTP server", zap.String("address", s.server.Addr))
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.log.Error("HTTP server exited with error", zap.Error(err))
		return err
	}
	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.cfg.ShutdownTimeoutSeconds)*time.Second)
	defer cancel()

	s.log.Info("Shutting down HTTP server")
	if err := s.server.Shutdown(ctx); err != nil {
		s.log.Error("Error during HTTP server shutdown", zap.Error(err))
		return err
	}

	s.log.Info("HTTP server shutdown successful")
	return nil
}
