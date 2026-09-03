package domain

import (
	"context"
	"net/mail"
	"regexp"

	"github.com/adm87/onyx-server/pkg/grpc"
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

// Credential is the stored record — no tokens, has the password hash instead.
type Credential struct {
	Subject      string
	Email        string
	PasswordHash string
}

type IdentityStore interface {
	SaveCredential(ctx context.Context, email string, password string) (*Credential, *grpc.Error)
	GetCredentialBySubject(ctx context.Context, subject string) (*Credential, *grpc.Error)
	GetCredentialByEmail(ctx context.Context, email string) (*Credential, *grpc.Error)
}

func ValidateCrededntials(email, password string) *grpc.Error {
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
