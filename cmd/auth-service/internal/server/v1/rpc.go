package v1

import (
	"context"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/errs/reasons"
	"github.com/adm87/onyx-server/pkg/config"
	g "github.com/adm87/onyx-server/pkg/grpc"
	authv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/auth/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

type AuthSvcRpc struct {
	authv1.UnimplementedAuthServiceServer

	cfg  *config.Config
	log  *zap.Logger
	errs *g.Errors
}

func NewAuthSvcRpc(cfg *config.Config, log *zap.Logger) *AuthSvcRpc {
	return &AuthSvcRpc{
		cfg:  cfg,
		log:  log,
		errs: g.NewErrors("v1.auth.svc"),
	}
}

func (s *AuthSvcRpc) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	s.log.Info("Register called", zap.Any("request", req))
	return nil, s.errs.New(codes.Unimplemented, reasons.Unimplemented,
		g.WithMessage("Auth v1 Register not implemented"),
	)
}

func (s *AuthSvcRpc) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	s.log.Info("Login called", zap.Any("request", req))
	return nil, s.errs.New(codes.Unimplemented, reasons.Unimplemented,
		g.WithMessage("Auth v1 Login not implemented"),
	)
}
