package v1

import (
	"github.com/adm87/onyx-server/pkg/config"
	authv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/auth/v1"
	"go.uber.org/zap"
)

type AuthSvcRpc struct {
	authv1.UnimplementedAuthServiceServer

	cfg *config.Config
	log *zap.Logger
}

func NewAuthSvcRpc(cfg *config.Config, log *zap.Logger) *AuthSvcRpc {
	return &AuthSvcRpc{
		cfg: cfg,
		log: log,
	}
}
