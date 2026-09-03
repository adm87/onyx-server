package v1

import (
	"context"

	"github.com/adm87/onyx-server/pkg/config"
	userv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/user/v1"
	"go.uber.org/zap"
)

type UserService struct {
	userv1.UnimplementedUserServiceServer

	cfg *config.Config
	log *zap.Logger
}

func NewUserService(cfg *config.Config, log *zap.Logger) *UserService {
	return &UserService{
		cfg: cfg,
		log: log,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	return nil, nil
}
