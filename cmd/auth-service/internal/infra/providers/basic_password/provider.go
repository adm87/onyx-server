package basicpassword

import (
	"context"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/domain"
	"github.com/adm87/onyx-server/pkg/config"
	"go.uber.org/zap"
)

type Authenticator struct {
	cfg   *config.Config
	log   *zap.Logger
	store domain.IdentityStore
}

func NewAuthenticator(cfg *config.Config, log *zap.Logger, store domain.IdentityStore) *Authenticator {
	return &Authenticator{
		cfg:   cfg,
		log:   log,
		store: store,
	}
}

func (a *Authenticator) Register(ctx context.Context, credentials domain.Credentials) (*domain.Identity, error) {
	// Implementation for registering a new user with basic password authentication
	return nil, nil
}

func (a *Authenticator) Authenticate(ctx context.Context, credentials domain.Credentials) (*domain.Identity, error) {
	// Implementation for authenticating a user with basic password authentication
	return nil, nil
}
