package memory

import (
	"context"
	"sync"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/domain"
)

// IdentityStore is an in-memory implementation of the IdentityRepository interface.
// Use this for testing or development purposes.
type IdentityStore struct {
	mu      sync.RWMutex
	byID    map[string]*domain.Credential
	byEmail map[string]*domain.Credential
}

func NewIdentityStore() *IdentityStore {
	return &IdentityStore{
		byID:    make(map[string]*domain.Credential),
		byEmail: make(map[string]*domain.Credential),
	}
}

func (s *IdentityStore) SaveCredential(ctx context.Context, credential *domain.Credential) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.byEmail[credential.Email]; exists {
		return domain.ErrEmailTaken
	}

	s.byID[credential.Subject] = credential
	s.byEmail[credential.Email] = credential
	return nil
}

func (s *IdentityStore) GetCredentialBySubject(ctx context.Context, subject string) (*domain.Credential, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cred, ok := s.byID[subject]
	if !ok {
		return nil, domain.ErrCredentialNotFound
	}
	return cred, nil
}

func (s *IdentityStore) GetCredentialByEmail(ctx context.Context, email string) (*domain.Credential, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cred, ok := s.byEmail[email]
	if !ok {
		return nil, domain.ErrCredentialNotFound
	}
	return cred, nil
}
