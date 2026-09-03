package v1

import (
	"context"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/domain"
	"github.com/adm87/onyx-server/pkg/config"
	"github.com/adm87/onyx-server/pkg/grpc"
	authv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/auth/v1"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
)

type AuthService struct {
	authv1.UnimplementedAuthServiceServer

	cfg  *config.Config
	log  *zap.Logger
	errs *grpc.ErrorProducer

	store domain.IdentityStore
}

func NewAuthService(cfg *config.Config, log *zap.Logger, store domain.IdentityStore) *AuthService {
	return &AuthService{
		cfg:   cfg,
		log:   log,
		store: store,
		errs:  grpc.NewErrorProducer("auth.v1", log),
	}
}

func (s *AuthService) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	email, password := req.GetEmail(), req.GetPassword()
	if gErr := domain.ValidateCrededntials(email, password); gErr != nil {
		return nil, s.errs.From(gErr)
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, s.errs.New(
			grpc.WithErrorCode(codes.Internal),
			grpc.WithErrorReason(domain.ReasonInternal),
		)
	}

	creds, gErr := s.store.SaveCredential(ctx, email, string(passHash))
	if gErr != nil {
		return nil, s.errs.From(gErr)
	}
	_ = creds
	return nil, nil
}

func (s *AuthService) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	s.log.Info("Login called", zap.Any("request", req))
	return nil, nil
}
