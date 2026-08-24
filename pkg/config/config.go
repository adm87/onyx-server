package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Gateway  GatewayConfig  `envPrefix:"GATEWAY_"`
	Logger   LoggerConfig   `envPrefix:"LOGGER_"`
	Auth     AuthConfig     `envPrefix:"AUTH_"`
	User     UserConfig     `envPrefix:"USER_"`
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
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
	Name            string     `env:"NAME" envDefault:"gateway"`
	Http            HttpConfig `envPrefix:"HTTP_"`
	EnableSwaggerUI bool       `env:"ENABLE_SWAGGER_UI" envDefault:"false"`
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

type PostgresConfig struct {
	Host                   string `env:"HOST" envDefault:"localhost"`
	Port                   int    `env:"PORT" envDefault:"5432"`
	User                   string `env:"USER" envDefault:"postgres"`
	Password               string `env:"PASSWORD" envDefault:"password"`
	DBName                 string `env:"DB_NAME" envDefault:"onyx"`
	SSLMode                string `env:"SSL_MODE" envDefault:"disable"`
	MaxOpenConns           int    `env:"MAX_OPEN_CONNS" envDefault:"25"`
	MaxIdleConns           int    `env:"MAX_IDLE_CONNS" envDefault:"25"`
	ConnMaxLifetimeSeconds int    `env:"CONN_MAX_LIFETIME_SECONDS" envDefault:"300"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
