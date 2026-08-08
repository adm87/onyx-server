package handler

import (
	"context"

	"github.com/adm87/onyx-server/internal/gateway/client"
	gateway "github.com/adm87/onyx-server/internal/gateway/gen"
	"github.com/adm87/onyx-server/internal/gateway/handler/auth"
	"github.com/adm87/onyx-server/internal/gateway/handler/user"
	"go.uber.org/zap"
)

type Handler struct {
	*auth.AuthHandler
	*user.UserHandler
}

var _ gateway.StrictServerInterface = (*Handler)(nil)

func NewHandler(clients *client.Clients, log *zap.Logger) *Handler {
	return &Handler{
		AuthHandler: auth.NewAuthHandler(clients.Auth, log),
		UserHandler: user.NewUserHandler(clients.User, log),
	}
}

func (h *Handler) Healthz(ctx context.Context, req gateway.HealthzRequestObject) (gateway.HealthzResponseObject, error) {
	return gateway.Healthz200Response{}, nil
}
