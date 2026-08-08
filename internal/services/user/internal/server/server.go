package server

import (
	"github.com/adm87/onyx-server/internal/config"
	userv1 "github.com/adm87/onyx-server/proto/gen/user/v1"
	"go.uber.org/zap"
)

type UserServer struct {
	userv1.UnimplementedUserServiceServer

	cfg *config.Config
	log *zap.Logger
}

func New(cfg *config.Config, log *zap.Logger) *UserServer {
	return &UserServer{
		cfg: cfg,
		log: log,
	}
}
