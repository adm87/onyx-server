package server

import (
	"context"

	authv1 "github.com/adm87/onyx-server/proto/gen/auth/v1"
)

func (s *AuthServer) RefreshToken(ctx context.Context, req *authv1.RefreshTokenRequest) (*authv1.TokenResponse, error) {
	return &authv1.TokenResponse{}, nil
}
