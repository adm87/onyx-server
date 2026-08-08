package server

import (
	"context"

	authv1 "github.com/adm87/onyx-server/proto/gen/auth/v1"
)

func (s *AuthServer) ValidateToken(ctx context.Context, req *authv1.ValidateTokenRequest) (*authv1.ValidateTokenResponse, error) {
	return &authv1.ValidateTokenResponse{}, nil
}
