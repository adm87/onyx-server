package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/adm87/onyx-server/internal/services/user/internal/domain"
	"github.com/adm87/onyx-server/internal/services/user/internal/repository"
)

type UserRepository struct {
	db *sql.DB
}

var _ repository.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	const query = `
		INSERT INTO users (id, email, display_name)
		VALUES ($1, $2, $3)
		RETURNING created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		user.ID, user.Email, user.DisplayName,
	).Scan(&user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	const query = `
		SELECT id, email, display_name, created_at, updated_at
		FROM users WHERE id = $1`
	return r.scanOne(r.db.QueryRowContext(ctx, query, id))
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	const query = `
		SELECT id, email, display_name, created_at, updated_at
		FROM users WHERE email = $1`
	return r.scanOne(r.db.QueryRowContext(ctx, query, email))
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	const query = `
		UPDATE users SET display_name = $1
		WHERE id = $2
		RETURNING updated_at`
	err := r.db.QueryRowContext(ctx, query, user.DisplayName, user.ID).Scan(&user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return repository.ErrNotFound
	}
	return err
}

func (r *UserRepository) scanOne(row *sql.Row) (*domain.User, error) {
	var u domain.User
	err := row.Scan(&u.ID, &u.Email, &u.DisplayName, &u.CreatedAt, &u.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan user: %w", err)
	}
	return &u, nil
}
