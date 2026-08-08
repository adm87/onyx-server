package domain

import "time"

type Identity struct {
	ID           string
	Email        string
	PasswordHash string
	AuthProvider string
	MFAEnabled   bool
	TOTPSecret   *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
