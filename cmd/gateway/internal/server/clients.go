package server

import (
	"context"
	"errors"

	"github.com/adm87/onyx-server/pkg/config"
	g "github.com/adm87/onyx-server/pkg/grpc"
	authv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/auth/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type svcClients map[string]*g.Client

func (sc svcClients) Close() error {
	var errs []error
	for _, client := range sc {
		errs = append(errs, client.Close())
	}
	return errors.Join(errs...)
}

func createSvcClients(ctx context.Context, cfg *config.Config, log *zap.Logger, mux *runtime.ServeMux) (svcClients, error) {
	authClient := g.NewClient(cfg.Auth.Svc.Name, &cfg.Auth.Svc.Grpc, log)
	if err := authClient.Connect(grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		log.Error("Failed to connect to Auth service", zap.Error(err))
		return nil, err
	}

	if err := authv1.RegisterAuthServiceHandler(ctx, mux, authClient.Conn()); err != nil {
		log.Error("Failed to register Auth service handler", zap.Error(err))
		return nil, err
	}

	return svcClients{
		cfg.Auth.Svc.Name: authClient,
	}, nil
}
