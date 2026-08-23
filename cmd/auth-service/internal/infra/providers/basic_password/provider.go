package basicpassword

import (
	"context"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/domain"
)

type Authenticator struct {
	repo domain.IdentityStore
}

func NewAuthenticator(repo domain.IdentityStore) *Authenticator {
	return &Authenticator{
		repo: repo,
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
