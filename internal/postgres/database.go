package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adm87/onyx-server/internal/config"
	"github.com/adm87/onyx-server/pkg/database"
	"go.uber.org/zap"
)

type DB struct {
	cfg *config.PostgresConfig
	sql *sql.DB
	log *zap.Logger
}

var _ database.Database = (*DB)(nil)

func NewDB(cfg *config.PostgresConfig, log *zap.Logger) *DB {
	return &DB{
		cfg: cfg,
		log: log,
	}
}

func (db *DB) Connect(ctx context.Context) error {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s,public",
		db.cfg.Host, db.cfg.Port, db.cfg.User, db.cfg.Password, db.cfg.DBName, db.cfg.Schema,
	)

	sqlDB, err := sql.Open("pgx", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	sqlDB.SetMaxOpenConns(db.cfg.OpenConns)
	sqlDB.SetMaxIdleConns(db.cfg.IdleConns)

	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	db.sql = sqlDB
	db.log.Info("connected to postgres", zap.String("schema", db.cfg.Schema))
	return nil
}

func (db *DB) Close(_ context.Context) error {
	if db.sql == nil {
		return nil
	}
	if err := db.sql.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	return nil
}

func (db *DB) Ping(ctx context.Context) error {
	if err := db.sql.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	return nil
}

func (db *DB) SQL() *sql.DB {
	return db.sql
}
