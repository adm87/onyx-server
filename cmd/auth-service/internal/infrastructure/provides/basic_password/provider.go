package basicpassword

import (
	"context"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/domain"
)

type Provider struct {
	repo domain.IdentityStore
}

func NewProvider(repo domain.IdentityStore) *Provider {
	return &Provider{
		repo: repo,
	}
}

func (p *Provider) Register(ctx context.Context, credentials domain.Credentials) (*domain.Identity, error) {
	// Implementation for registering a new user with basic password authentication
	return nil, nil
}

func (p *Provider) Authenticate(ctx context.Context, credentials domain.Credentials) (*domain.Identity, error) {
	// Implementation for authenticating a user with basic password authentication
	return nil, nil
}
