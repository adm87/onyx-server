package domain

import (
	"context"
	"net/mail"
	"regexp"

	"github.com/adm87/onyx-server/pkg/server/grpc"
	"google.golang.org/grpc/codes"
)

var (
	reLetter = regexp.MustCompile(`[A-Za-z]`)
	reDigit  = regexp.MustCompile(`\d`)
)

type StoreType string

const (
	StoreTypeInMemory StoreType = "inmemory"
	StoreTypePostgres StoreType = "postgres"
)

func (s StoreType) IsValid() bool {
	switch s {
	case StoreTypeInMemory, StoreTypePostgres:
		return true
	default:
		return false
	}
}

type Identity struct {
	Subject      string
	Email        string
	PasswordHash string
}

type IdentityStore interface {
	CreateIdentity(ctx context.Context, email string, password string) (*Identity, error)
	GetIdentityBySubject(ctx context.Context, subject string) (*Identity, error)
	GetIdentityByEmail(ctx context.Context, email string) (*Identity, error)
}

type TokenProvider interface {
	GenerateToken(ctx context.Context, subject string) (string, error)
}

func ValidateCrededntials(email, password string) error {
	if email == "" {
		return &grpc.Error{
			Code:    codes.InvalidArgument,
			Reason:  ReasonInvalidCredentials,
			Message: "email is required",
		}
	}
	if password == "" {
		return &grpc.Error{
			Code:    codes.InvalidArgument,
			Reason:  ReasonInvalidCredentials,
			Message: "password is required",
		}
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return &grpc.Error{
			Code:    codes.InvalidArgument,
			Reason:  ReasonInvalidCredentials,
			Message: "invalid email",
		}
	}
	if !isValidPassword(password) {
		return &grpc.Error{
			Code:    codes.InvalidArgument,
			Reason:  ReasonInvalidCredentials,
			Message: "password must be at least 8 characters and contain a letter and a digit",
		}
	}
	return nil
}

func isValidPassword(pw string) bool {
	return len(pw) >= 8 &&
		reLetter.MatchString(pw) &&
		reDigit.MatchString(pw)
}
