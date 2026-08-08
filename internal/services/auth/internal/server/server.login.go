package server

import (
	"context"

	authv1 "github.com/adm87/onyx-server/proto/gen/auth/v1"
)

func (s *AuthServer) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.TokenResponse, error) {
	return &authv1.TokenResponse{}, nil
}
