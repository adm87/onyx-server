package gateway

import (
	"fmt"

	"github.com/adm87/onyx-server/internal/config"
	"github.com/adm87/onyx-server/internal/gateway/client"
	gateway "github.com/adm87/onyx-server/internal/gateway/gen"
	"github.com/adm87/onyx-server/internal/gateway/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Run(cfg *config.Config, log *zap.Logger) error {
	if err := validateConfig(cfg); err != nil {
		return err
	}

	clients, err := client.New(
		cfg.Services.Auth.GRPC.Addr,
		cfg.Services.User.GRPC.Addr,
	)
	if err != nil {
		return err
	}

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	h := handler.NewHandler(clients, log)
	gateway.RegisterHandlers(e, gateway.NewStrictHandler(h, nil))

	log.Info("Starting gateway server", zap.String("addr", cfg.Gateway.HTTP.Port))
	return e.Start(fmt.Sprintf(":%s", cfg.Gateway.HTTP.Port))
}

func validateConfig(cfg *config.Config) error {
	if cfg.Gateway.HTTP.Port == "" {
		return fmt.Errorf("gateway http port is not set")
	}

	if cfg.Services.Auth.GRPC.Addr == "" {
		return fmt.Errorf("auth service grpc addr is not set")
	}

	if cfg.Services.User.GRPC.Addr == "" {
		return fmt.Errorf("user service grpc addr is not set")
	}

	return nil
}
