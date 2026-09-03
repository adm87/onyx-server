package inmemory

import (
	"context"
	"sync"
	"uuid"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/domain"
	"github.com/adm87/onyx-server/pkg/config"
	"github.com/adm87/onyx-server/pkg/server/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

type InMemoryIdentityStore struct {
	cfg *config.Config
	log *zap.Logger

	mu      sync.RWMutex
	byID    map[string]*domain.Identity
	byEmail map[string]*domain.Identity
}

func NewInMemoryIdentityStore(cfg *config.Config, log *zap.Logger) *InMemoryIdentityStore {
	return &InMemoryIdentityStore{
		cfg:     cfg,
		log:     log,
		byID:    make(map[string]*domain.Identity),
		byEmail: make(map[string]*domain.Identity),
	}
}

func (s *InMemoryIdentityStore) CreateIdentity(ctx context.Context, email string, password string) (*domain.Identity, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.byEmail[email]; exists {
		return nil, &grpc.Error{
			Code:    codes.AlreadyExists,
			Reason:  domain.ReasonEmailUnavailable,
			Message: "email is already in use",
		}
	}

	subject := uuid.New().String()
	creds := &domain.Identity{
		Subject:      subject,
		Email:        email,
		PasswordHash: password,
	}

	s.byID[subject] = creds
	s.byEmail[email] = creds

	return creds, nil
}

func (s *InMemoryIdentityStore) GetIdentityBySubject(ctx context.Context, subject string) (*domain.Identity, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	credential, exists := s.byID[subject]
	if !exists {
		return nil, &grpc.Error{
			Code:    codes.NotFound,
			Reason:  domain.ReasonSubjectNotFound,
			Message: "subject not found",
		}
	}

	return credential, nil
}

func (s *InMemoryIdentityStore) GetIdentityByEmail(ctx context.Context, email string) (*domain.Identity, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	credential, exists := s.byEmail[email]
	if !exists {
		return nil, &grpc.Error{
			Code:    codes.NotFound,
			Reason:  domain.ReasonEmailNotFound,
			Message: "email not found",
		}
	}

	return credential, nil
}

func (s *InMemoryIdentityStore) Connect() error { return nil }
func (s *InMemoryIdentityStore) Close() error   { return nil }
func (s *InMemoryIdentityStore) Ping() error    { return nil }
