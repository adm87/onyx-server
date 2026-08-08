package user

import (
	"context"

	gateway "github.com/adm87/onyx-server/internal/gateway/gen"
	userv1 "github.com/adm87/onyx-server/proto/gen/user/v1"
	"go.uber.org/zap"
)

type UserHandler struct {
	userClient userv1.UserServiceClient
	log        *zap.Logger
}

func NewUserHandler(userClient userv1.UserServiceClient, log *zap.Logger) *UserHandler {
	return &UserHandler{
		userClient: userClient,
		log:        log,
	}
}

func (h *UserHandler) GetUser(ctx context.Context, req gateway.GetUserRequestObject) (gateway.GetUserResponseObject, error) {
	return gateway.GetUser204Response{}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req gateway.UpdateUserRequestObject) (gateway.UpdateUserResponseObject, error) {
	return gateway.UpdateUser204Response{}, nil
}
