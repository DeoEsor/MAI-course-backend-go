package config

import (
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
	"github.com/samber/do/v2"
)

type Config struct {
	Environment    string         `env:"ENVIRONMENT"`
	DatabaseConfig DatabaseConfig `envPrefix:"DATABASE_"`
}

func Load() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg *Config) Register(injector do.Injector) (*Config, error) {
	return cfg, nil
}
