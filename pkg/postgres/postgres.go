package postgres

import (
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/adm87/onyx-server/pkg/config"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	cfg *config.PostgresConfig
	db  *sql.DB
}

func NewPostgresConn(cfg *config.PostgresConfig, log *zap.Logger) (*PostgresDB, error) {
	log.Info("Connecting to PostgreSQL database",
		zap.String("host", cfg.Host),
		zap.Int("port", cfg.Port),
		zap.String("user", cfg.User),
		zap.String("dbname", cfg.DBName),
		zap.String("schema", cfg.Schema),
		zap.String("sslmode", cfg.SSLMode),
	)

	connStr := buildConnString(cfg)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetimeSeconds) * time.Second)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{
		cfg: cfg,
		db:  db,
	}, nil
}

func (p *PostgresDB) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

func (p *PostgresDB) Ping() error {
	if p.db != nil {
		return p.db.Ping()
	}
	return fmt.Errorf("database connection is not initialized")
}

func (p *PostgresDB) DB() *sql.DB {
	return p.db
}

func buildConnString(cfg *config.PostgresConfig) string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Path:   fmt.Sprintf("/%s", cfg.DBName),
	}

	q := u.Query()
	q.Set("sslmode", cfg.SSLMode)

	if cfg.Schema != "" {
		q.Set("search_path", cfg.Schema)
	}

	u.RawQuery = q.Encode()
	return u.String()
}
