package store

import (
	"fmt"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/domain"
	"github.com/adm87/onyx-server/cmd/auth-service/internal/store/inmemory"
	"github.com/adm87/onyx-server/cmd/auth-service/internal/store/postgres"
	"github.com/adm87/onyx-server/pkg/config"
	"go.uber.org/zap"
)

type IdentityStore interface {
	domain.IdentityStore

	Connect() error
	Close() error
	Ping() error
}

func NewIdentityStore(cfg *config.Config, log *zap.Logger) (IdentityStore, error) {
	log.Info("Creating identity store", zap.String("store_type", cfg.Auth.StoreType))
	switch domain.StoreType(cfg.Auth.StoreType) {
	case domain.StoreTypeInMemory:
		return inmemory.NewInMemoryIdentityStore(cfg, log), nil
	case domain.StoreTypePostgres:
		return postgres.NewPostgresIdentityStore(cfg, log)
	default:
		log.Error("Unsupported store type", zap.String("store_type", cfg.Auth.StoreType))
		return nil, fmt.Errorf("unsupported store type: %s", cfg.Auth.StoreType)
	}
}
