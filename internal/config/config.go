package config

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Environment string `env:"ENVIRONMENT"`
}

func (cfg *Config) Load() error {
	return env.Parse(cfg)
}
