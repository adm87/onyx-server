package grpc

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

type ServerShutdownHook func() error

type RunOpts struct {
	BeforeServerShutdown []ServerShutdownHook
	AfterServerShutdown  []ServerShutdownHook
}

type RunOption func(*RunOpts)

func defaultRunOpts() *RunOpts {
	return &RunOpts{
		BeforeServerShutdown: make([]ServerShutdownHook, 0),
		AfterServerShutdown:  make([]ServerShutdownHook, 0),
	}
}

func WithBeforeServerShutdownHooks(hooks ...ServerShutdownHook) RunOption {
	return func(opts *RunOpts) {
		opts.BeforeServerShutdown = append(opts.BeforeServerShutdown, hooks...)
	}
}

func WithAfterServerShutdownHooks(hooks ...ServerShutdownHook) RunOption {
	return func(opts *RunOpts) {
		opts.AfterServerShutdown = append(opts.AfterServerShutdown, hooks...)
	}
}

func Run(server *Server, opts ...RunOption) error {
	options := defaultRunOpts()
	for _, opt := range opts {
		opt(options)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.start()
	}()

	select {
	case <-signals:
		server.log.Info("Received shutdown signal, exiting...")

		var shutdownErrs []error

		// Execute before shutdown hooks
		if err := callShutdownHooks(options.BeforeServerShutdown); err != nil {
			shutdownErrs = append(shutdownErrs, err)
		}

		// Shutdown the gRPC server
		if err := server.shutdown(); err != nil {
			shutdownErrs = append(shutdownErrs, err)
		}

		// Execute after shutdown hooks
		if err := callShutdownHooks(options.AfterServerShutdown); err != nil {
			shutdownErrs = append(shutdownErrs, err)
		}

		return errors.Join(shutdownErrs...)

	case err := <-errCh:
		server.log.Error("gRPC server error", zap.Error(err))
		return err
	}
}

func callShutdownHooks(hooks []ServerShutdownHook) error {
	var errs []error
	for _, hook := range hooks {
		errs = append(errs, hook())
	}
	return errors.Join(errs...)
}
