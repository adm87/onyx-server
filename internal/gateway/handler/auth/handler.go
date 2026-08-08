package auth

import (
	"context"

	gateway "github.com/adm87/onyx-server/internal/gateway/gen"
	authv1 "github.com/adm87/onyx-server/proto/gen/auth/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AuthHandler struct {
	authClient authv1.AuthServiceClient
	log        *zap.Logger
}

func NewAuthHandler(authClient authv1.AuthServiceClient, log *zap.Logger) *AuthHandler {
	return &AuthHandler{
		authClient: authClient,
		log:        log,
	}
}

func (h *AuthHandler) Login(ctx context.Context, req gateway.LoginRequestObject) (gateway.LoginResponseObject, error) {
	resp, err := h.authClient.Login(ctx, &authv1.LoginRequest{}, grpc.WaitForReady(true))
	if err != nil {
		h.log.Error("failed to login", zap.Error(err))
		return nil, err
	}
	h.log.Info("login successful", zap.Any("response", resp))
	return gateway.Login204Response{}, nil
}

func (h *AuthHandler) Signup(ctx context.Context, req gateway.SignupRequestObject) (gateway.SignupResponseObject, error) {
	return gateway.Signup204Response{}, nil
}

func (h *AuthHandler) RefreshToken(ctx context.Context, req gateway.RefreshTokenRequestObject) (gateway.RefreshTokenResponseObject, error) {
	return gateway.RefreshToken204Response{}, nil
}
