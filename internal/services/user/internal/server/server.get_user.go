package server

import (
	"context"

	userv1 "github.com/adm87/onyx-server/proto/gen/user/v1"
)

func (s *UserServer) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.User, error) {
	return &userv1.User{}, nil
}
