package config

import (
	//Standard library packages
	"fmt"

	//External packages
	"github.com/caarlos0/env/v6"
)

type Config struct {
	DB DB `envPrefix:"DB_"`
}

type DB struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Name     string `env:"NAME" envDefault:"main"`
	SSLMode  string `env:"SSL_MODE" envDefault:"disable"`
	PoolSize int    `env:"POOL_SIZE" envDefault:"10"`
}

func Read(cfg *Config) error {
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	return nil
}
