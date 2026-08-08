package repository

import (
	"context"
	"errors"

	"github.com/adm87/onyx-server/internal/services/auth/internal/domain"
)

var ErrNotFound = errors.New("identity not found")

type IdentityRepository interface {
	Create(ctx context.Context, identity *domain.Identity) error
	FindByEmail(ctx context.Context, email string) (*domain.Identity, error)
	FindByID(ctx context.Context, id string) (*domain.Identity, error)
}
