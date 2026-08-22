package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Gateway HttpConfig   `envPrefix:"GATEWAY_"`
	Logger  LoggerConfig `envPrefix:"LOGGER_"`
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

func Load() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
