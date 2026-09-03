package grpc

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Reason string
type Metadata map[string]string

type Error struct {
	Code     codes.Code
	Reason   Reason
	Message  string
	Metadata Metadata
	Err      error
}

func (e *Error) Error() string {
	return fmt.Sprintf(
		"grpc error: code=%v reason=%v message=%q metadata=%v underlying=%v",
		e.Code, e.Reason, e.Message, e.Metadata, e.Err,
	)
}

type ErrorProducer struct {
	log    *zap.Logger
	domain string
}

func NewErrorProducer(domain string, log *zap.Logger) *ErrorProducer {
	return &ErrorProducer{
		log:    log,
		domain: domain,
	}
}

func (p *ErrorProducer) New(opts ...func(*Error)) error {
	e := &Error{}
	for _, opt := range opts {
		opt(e)
	}
	return p.From(e)
}

func (p *ErrorProducer) From(err error) error {
	gErr, ok := err.(*Error)
	if !ok {
		return p.New(
			WithErrorCode(codes.Internal),
			WithErrorReason("INTERNAL"),
			WithErrorMessage(err.Error()),
			WithError(err),
		)
	}

	if gErr.Metadata == nil {
		gErr.Metadata = Metadata{}
	}

	info := &errdetails.ErrorInfo{
		Domain:   p.domain,
		Reason:   string(gErr.Reason),
		Metadata: gErr.Metadata,
	}

	st, detailErr := status.New(gErr.Code, gErr.Message).WithDetails(info)
	if detailErr != nil {
		return errors.Join(
			fmt.Errorf("failed to attach error details: %w", detailErr),
			gErr.Err,
		)
	}

	statusErr := st.Err()

	p.log.Error(
		"grpc error emitted",
		zap.Error(statusErr),
		zap.String("reason", string(gErr.Reason)),
		zap.Any("metadata", gErr.Metadata),
	)

	return statusErr
}

func WithErrorCode(code codes.Code) func(*Error) {
	return func(e *Error) { e.Code = code }
}

func WithErrorReason(reason Reason) func(*Error) {
	return func(e *Error) { e.Reason = reason }
}

func WithErrorMessage(msg string) func(*Error) {
	return func(e *Error) { e.Message = msg }
}

func WithErrorMetadata(md Metadata) func(*Error) {
	return func(e *Error) { e.Metadata = md }
}

func WithError(err error) func(*Error) {
	return func(e *Error) { e.Err = err }
}
