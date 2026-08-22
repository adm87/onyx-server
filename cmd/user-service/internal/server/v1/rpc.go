package v1

import (
	"context"

	"github.com/adm87/onyx-server/cmd/user-service/errs/reasons"
	"github.com/adm87/onyx-server/pkg/config"
	g "github.com/adm87/onyx-server/pkg/grpc"
	userv1 "github.com/adm87/onyx-server/pkg/proto/gen/grpc/user/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

type UserSvcRpc struct {
	userv1.UnimplementedUserServiceServer

	cfg  *config.Config
	log  *zap.Logger
	errs *g.Errors
}

func NewUserSvcRpc(cfg *config.Config, log *zap.Logger) *UserSvcRpc {
	return &UserSvcRpc{
		cfg:  cfg,
		log:  log,
		errs: g.NewErrors("v1.user.svc"),
	}
}

func (s *UserSvcRpc) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	return nil, s.errs.New(codes.Unimplemented, reasons.Unimplemented,
		g.WithMessage("CreateUser is not implemented yet"),
	)
}
