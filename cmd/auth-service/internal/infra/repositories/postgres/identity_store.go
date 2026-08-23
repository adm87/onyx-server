package postgres

import (
	"context"
	"database/sql"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/domain"
	"github.com/adm87/onyx-server/pkg/config"
	"go.uber.org/zap"
)

type IdentityStore struct {
	cfg *config.PostgresConfig
	log *zap.Logger
	db  *sql.DB
}

func NewIdentityStore(cfg *config.PostgresConfig, log *zap.Logger, db *sql.DB) *IdentityStore {
	return &IdentityStore{
		cfg: cfg,
		log: log,
		db:  db,
	}
}

func (s *IdentityStore) SaveCredential(ctx context.Context, credential *domain.Credential) error {
	// Implementation to save the credential in the database
	return nil
}

func (s *IdentityStore) GetCredentialBySubject(ctx context.Context, subject string) (*domain.Credential, error) {
	// Implementation to retrieve the credential by subject from the database
	return nil, nil
}

func (s *IdentityStore) GetCredentialByEmail(ctx context.Context, email string) (*domain.Credential, error) {
	// Implementation to retrieve the credential by email from the database
	return nil, nil
}
