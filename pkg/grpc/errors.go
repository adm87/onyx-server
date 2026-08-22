package grpc

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Reason string

type Metadata map[string]string

type opts struct {
	Message  *string
	Metadata Metadata
}

type ErrOpt func(*opts)

func defaultOpts() *opts {
	return &opts{}
}

func WithMessage(message string) ErrOpt {
	return func(o *opts) {
		o.Message = &message
	}
}

func WithMetadata(metadata Metadata) ErrOpt {
	return func(o *opts) {
		o.Metadata = metadata
	}
}

// Errors is a factory helper for creating structured gRPC errors with a specific domain.
type Errors struct {
	domain string
}

func NewErrors(domain string) *Errors {
	return &Errors{
		domain: domain,
	}
}

func (e *Errors) New(code codes.Code, reason Reason, opts ...ErrOpt) error {
	options := defaultOpts()
	for _, opt := range opts {
		opt(options)
	}

	info := &errdetails.ErrorInfo{
		Domain: e.domain,
		Reason: string(reason),
	}

	if len(options.Metadata) > 0 {
		info.Metadata = options.Metadata
	}

	message := string(reason)
	if options.Message != nil {
		message = *options.Message
	}

	s, err := status.New(code, message).WithDetails(info)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to add error details: %v", err)
	}

	return s.Err()
}
