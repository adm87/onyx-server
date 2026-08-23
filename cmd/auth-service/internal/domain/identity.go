package domain

import "context"

// Identity is returned after a successful register or login. Never persisted directly —
// it carries tokens, which are ephemeral, not part of the stored record.
type Identity struct {
	Subject      string
	Email        string
	AccessToken  string
	RefreshToken string
}

// Credentials is what a caller submits to register or log in.
type Credentials struct {
	Email    string
	Password string
}

// Credential is the stored record — no tokens, has the password hash instead.
type Credential struct {
	Subject      string
	Email        string
	PasswordHash string
}

type IdentityProvider interface {
	Register(ctx context.Context, credentials Credentials) (*Identity, error)
	Authenticate(ctx context.Context, credentials Credentials) (*Identity, error)
}

type IdentityStore interface {
	SaveCredential(ctx context.Context, credential *Credential) error
	GetCredentialBySubject(ctx context.Context, subject string) (*Credential, error)
	GetCredentialByEmail(ctx context.Context, email string) (*Credential, error)
}
