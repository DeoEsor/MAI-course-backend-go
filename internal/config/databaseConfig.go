package config

import "time"

type DatabaseConfig struct {
	DSN                       string        `env:"DSN" env-required:"true"`
	DatabaseMaxConnections    int           `env:"MAX_CONNECTIONS" envDefault:"5"`
	DatabaseConnectionTimeout time.Duration `env:"CONNECTION_TIMEOUT" envDefault:"5s"`
}
