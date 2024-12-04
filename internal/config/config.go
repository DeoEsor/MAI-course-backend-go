package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/samber/do/v2"
	log "github.com/sirupsen/logrus"

	_ "github.com/joho/godotenv/autoload"
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
	log.
		WithFields(log.Fields{
			"service": "config",
		}).
		Info("service invoked")
	return cfg, nil
}
