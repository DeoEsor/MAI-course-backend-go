package config

import (
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Environment string `env:"ENVIRONMENT"`
}

func (config *Config) Load() error {
	return env.Parse(config)
}
