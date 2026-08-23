package domain

import "errors"

const (
	ReasonUnimplemented        = "UNIMPLEMENTED"
	ReasonRegistrationFailed   = "REGISTRATION_FAILED"
	ReasonAuthenticationFailed = "AUTHENTICATION_FAILED"
)

var (
	ErrEmailTaken         = errors.New("email already taken")
	ErrCredentialNotFound = errors.New("credential not found")
)
