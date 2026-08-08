package client

import (
	"fmt"

	authv1 "github.com/adm87/onyx-server/proto/gen/auth/v1"
	userv1 "github.com/adm87/onyx-server/proto/gen/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Auth authv1.AuthServiceClient
	User userv1.UserServiceClient
}

func New(authAddr, userAddr string) (*Clients, error) {
	authConn, err := grpc.NewClient(authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("dial auth-service: %w", err)
	}

	userConn, err := grpc.NewClient(userAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("dial user-service: %w", err)
	}

	return &Clients{
		Auth: authv1.NewAuthServiceClient(authConn),
		User: userv1.NewUserServiceClient(userConn),
	}, nil
}
