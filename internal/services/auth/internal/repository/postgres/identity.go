package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/adm87/onyx-server/internal/services/auth/internal/domain"
	"github.com/adm87/onyx-server/internal/services/auth/internal/repository"
)

type IdentityRepository struct {
	db *sql.DB
}

var _ repository.IdentityRepository = (*IdentityRepository)(nil)

func NewIdentityRepository(db *sql.DB) *IdentityRepository {
	return &IdentityRepository{db: db}
}

func (r *IdentityRepository) Create(ctx context.Context, identity *domain.Identity) error {
	const query = `
		INSERT INTO identities (id, email, password_hash, auth_provider, mfa_enabled, totp_secret)
		VALUES (gen_random_uuid(), $1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		identity.Email, identity.PasswordHash, identity.AuthProvider, identity.MFAEnabled, identity.TOTPSecret,
	).Scan(&identity.ID, &identity.CreatedAt, &identity.UpdatedAt)
}

func (r *IdentityRepository) FindByEmail(ctx context.Context, email string) (*domain.Identity, error) {
	const query = `
		SELECT id, email, password_hash, auth_provider, mfa_enabled, totp_secret, created_at, updated_at
		FROM identities WHERE email = $1`
	return r.scanOne(r.db.QueryRowContext(ctx, query, email))
}

func (r *IdentityRepository) FindByID(ctx context.Context, id string) (*domain.Identity, error) {
	const query = `
		SELECT id, email, password_hash, auth_provider, mfa_enabled, totp_secret, created_at, updated_at
		FROM identities WHERE id = $1`
	return r.scanOne(r.db.QueryRowContext(ctx, query, id))
}

func (r *IdentityRepository) scanOne(row *sql.Row) (*domain.Identity, error) {
	var i domain.Identity
	err := row.Scan(&i.ID, &i.Email, &i.PasswordHash, &i.AuthProvider, &i.MFAEnabled, &i.TOTPSecret, &i.CreatedAt, &i.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan identity: %w", err)
	}
	return &i, nil
}
