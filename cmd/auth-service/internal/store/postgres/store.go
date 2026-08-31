package postgres

import (
	"context"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/domain"
	"github.com/adm87/onyx-server/pkg/config"
	"github.com/adm87/onyx-server/pkg/postgres"
	"go.uber.org/zap"
)

type PostgresIdentityStore struct {
	cfg  *config.Config
	log  *zap.Logger
	conn *postgres.PostgresDB
}

func NewPostgresIdentityStore(cfg *config.Config, log *zap.Logger) (*PostgresIdentityStore, error) {
	pgConn, err := postgres.NewPostgresConn(&cfg.Postgres, log)
	if err != nil {
		return nil, err
	}
	return &PostgresIdentityStore{
		cfg:  cfg,
		log:  log,
		conn: pgConn,
	}, nil
}

func (s *PostgresIdentityStore) Connect() error {
	return s.conn.Ping()
}

func (s *PostgresIdentityStore) Close() error {
	return s.conn.Close()
}

func (s *PostgresIdentityStore) Ping() error {
	return s.conn.Ping()
}

func (s *PostgresIdentityStore) SaveCredential(ctx context.Context, credential *domain.Credential) error {
	return nil
}

func (s *PostgresIdentityStore) GetCredentialBySubject(ctx context.Context, subject string) (*domain.Credential, error) {
	return nil, nil
}

func (s *PostgresIdentityStore) GetCredentialByEmail(ctx context.Context, email string) (*domain.Credential, error) {
	return nil, nil
}
