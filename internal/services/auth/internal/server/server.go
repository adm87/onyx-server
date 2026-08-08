package server

import (
	"github.com/adm87/onyx-server/internal/config"
	authv1 "github.com/adm87/onyx-server/proto/gen/auth/v1"
	"go.uber.org/zap"
)

type AuthServer struct {
	authv1.UnimplementedAuthServiceServer

	cfg *config.Config
	log *zap.Logger
}

func New(cfg *config.Config, log *zap.Logger) *AuthServer {
	return &AuthServer{
		cfg: cfg,
		log: log,
	}
}
