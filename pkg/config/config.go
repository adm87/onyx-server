package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Gateway HttpConfig `envPrefix:"GATEWAY_"`
}

type HttpConfig struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port int    `env:"PORT" envDefault:"8080"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
