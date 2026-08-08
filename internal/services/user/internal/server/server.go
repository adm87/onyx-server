package server

import (
	"github.com/adm87/onyx-server/internal/config"
	"github.com/adm87/onyx-server/internal/services/user/internal/repository/postgres"
	userv1 "github.com/adm87/onyx-server/proto/gen/user/v1"
	"go.uber.org/zap"
)

type UserServer struct {
	userv1.UnimplementedUserServiceServer

	cfg *config.Config
	log *zap.Logger

	userRepo *postgres.UserRepository
}

func New(cfg *config.Config, log *zap.Logger, userRepo *postgres.UserRepository) *UserServer {
	return &UserServer{
		cfg:      cfg,
		log:      log,
		userRepo: userRepo,
	}
}
