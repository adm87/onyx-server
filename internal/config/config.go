package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Gateway  GatewayConfig  `envPrefix:"GATEWAY_"`
	Logger   LoggerConfig   `envPrefix:"LOGGER_"`
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
	Services ServicesConfig `envPrefix:"SERVICES_"`
}

type LoggerConfig struct {
	Level  string `env:"LEVEL" envDefault:"info"`
	Prefix string `env:"PREFIX" envDefault:"onyx"`
}

type PostgresConfig struct {
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	DBName   string `env:"DB"`
}

type GatewayConfig struct {
	HTTP HTTPConfig `envPrefix:"HTTP_"`
}

type ServicesConfig struct {
	Auth AuthServiceConfig `envPrefix:"AUTH_"`
	User UserServiceConfig `envPrefix:"USER_"`
}

type AuthServiceConfig struct {
	GRPC GRPCConfig `envPrefix:"GRPC_"`
}

type UserServiceConfig struct {
	GRPC GRPCConfig `envPrefix:"GRPC_"`
}

type GRPCConfig struct {
	Port string `env:"PORT"`
	Addr string `env:"ADDR"`
}

type HTTPConfig struct {
	Port string `env:"PORT"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
