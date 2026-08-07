package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
}

type PostgresConfig struct {
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	DBName   string `env:"DB"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
