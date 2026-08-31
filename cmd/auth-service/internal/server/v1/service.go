package v1

import (
	"context"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/domain"
	"github.com/adm87/onyx-server/pkg/config"
	g "github.com/adm87/onyx-server/pkg/grpc"
	authv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/auth/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

type AuthService struct {
	authv1.UnimplementedAuthServiceServer

	cfg  *config.Config
	log  *zap.Logger
	errs *g.Errors

	store domain.IdentityStore
}

func NewAuthService(cfg *config.Config, log *zap.Logger, store domain.IdentityStore) *AuthService {
	return &AuthService{
		cfg:   cfg,
		log:   log,
		errs:  g.NewErrors("v1.auth.svc"),
		store: store,
	}
}

func (s *AuthService) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	return nil, s.errs.New(codes.Unimplemented, domain.ReasonUnimplemented,
		g.WithMessage("Auth v1 Register not implemented"),
	)
}

func (s *AuthService) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	s.log.Info("Login called", zap.Any("request", req))
	return nil, s.errs.New(codes.Unimplemented, domain.ReasonUnimplemented,
		g.WithMessage("Auth v1 Login not implemented"),
	)
}
