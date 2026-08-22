package v1

import (
	"github.com/adm87/onyx-server/pkg/config"
	userv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/user/v1"
	"go.uber.org/zap"
)

type UserSvcRpc struct {
	userv1.UnimplementedUserServiceServer

	cfg *config.Config
	log *zap.Logger
}

func NewUserSvcRpc(cfg *config.Config, log *zap.Logger) *UserSvcRpc {
	return &UserSvcRpc{
		cfg: cfg,
		log: log,
	}
}
