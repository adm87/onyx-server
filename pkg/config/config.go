package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Gateway GatewayConfig `envPrefix:"GATEWAY_"`
	Logger  LoggerConfig  `envPrefix:"LOGGER_"`
	Auth    AuthConfig    `envPrefix:"AUTH_"`
	User    UserConfig    `envPrefix:"USER_"`
}

type LoggerConfig struct {
	Level string `env:"LEVEL" envDefault:"info"`
}

type HttpConfig struct {
	Host                   string `env:"HOST" envDefault:"0.0.0.0"`
	Port                   int    `env:"PORT" envDefault:"8080"`
	IdleTimeoutSeconds     int    `env:"IDLE_TIMEOUT_SECONDS" envDefault:"60"`
	ReadTimeoutSeconds     int    `env:"READ_TIMEOUT_SECONDS" envDefault:"10"`
	WriteTimeoutSeconds    int    `env:"WRITE_TIMEOUT_SECONDS" envDefault:"10"`
	ShutdownTimeoutSeconds int    `env:"SHUTDOWN_TIMEOUT_SECONDS" envDefault:"10"`
}

type GrpcConfig struct {
	Host                   string `env:"HOST" envDefault:"0.0.0.0"`
	Port                   int    `env:"PORT" envDefault:"50051"`
	ConnectTimeoutSeconds  int    `env:"CONNECT_TIMEOUT_SECONDS" envDefault:"5"`
	ShutdownTimeoutSeconds int    `env:"SHUTDOWN_TIMEOUT_SECONDS" envDefault:"10"`
}

type GatewayConfig struct {
	Name string     `env:"NAME" envDefault:"gateway"`
	Http HttpConfig `envPrefix:"HTTP_"`
}

type SvcConfig struct {
	Name string     `env:"NAME" envDefault:""`
	Grpc GrpcConfig `envPrefix:"GRPC_"`
}

type AuthConfig struct {
	Svc SvcConfig `envPrefix:"SVC_"`
}

type UserConfig struct {
	Svc SvcConfig `envPrefix:"SVC_"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
