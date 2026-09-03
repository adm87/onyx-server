package domain

import "github.com/adm87/onyx-server/pkg/server/grpc"

const (
	ReasonInternal           grpc.Reason = "INTERNAL"
	ReasonInvalidCredentials grpc.Reason = "INVALID_CREDENTIALS"
	ReasonEmailUnavailable   grpc.Reason = "EMAIL_UNAVAILABLE"
	ReasonSubjectNotFound    grpc.Reason = "SUBJECT_NOT_FOUND"
	ReasonEmailNotFound      grpc.Reason = "EMAIL_NOT_FOUND"
)
