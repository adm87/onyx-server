package inmemory

import (
	"context"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/domain"
	"github.com/adm87/onyx-server/pkg/config"
	"go.uber.org/zap"
)

type InMemoryIdentityStore struct {
	cfg *config.Config
	log *zap.Logger
}

func NewInMemoryIdentityStore(cfg *config.Config, log *zap.Logger) *InMemoryIdentityStore {
	return &InMemoryIdentityStore{
		cfg: cfg,
		log: log,
	}
}

func (s *InMemoryIdentityStore) Connect() error {
	// In-memory store doesn't require a connection, so we just return nil
	return nil
}

func (s *InMemoryIdentityStore) Close() error {
	// In-memory store doesn't require closing, so we just return nil
	return nil
}

func (s *InMemoryIdentityStore) Ping() error {
	// In-memory store doesn't require pinging, so we just return nil
	return nil
}

func (s *InMemoryIdentityStore) SaveCredential(ctx context.Context, credential *domain.Credential) error {
	return nil
}

func (s *InMemoryIdentityStore) GetCredentialBySubject(ctx context.Context, subject string) (*domain.Credential, error) {
	return nil, nil
}

func (s *InMemoryIdentityStore) GetCredentialByEmail(ctx context.Context, email string) (*domain.Credential, error) {
	return nil, nil
}
