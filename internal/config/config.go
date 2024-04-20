package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	DB DB `envPrefix:"DB_"`
}

type DB struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	Name     string `env:"DB_NAME" envDefault:"main"`
	SSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
	PoolSize int    `env:"DB_POOL_SIZE" envDefault:"10"`
}

func Read(cfg *Config) error {
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	return nil
}
