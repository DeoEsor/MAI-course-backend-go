package main

import (
	"fmt"

	gobasics "github.com/DeoEsor/MAI-course-backend-go/internal/gobasics"

	_ "github.com/joho/godotenv/autoload"

	"github.com/caarlos0/env/v11"
)

func main() {
	var config Config
	err := config.Load()
	if err != nil {
		fmt.Printf("error while parsing config: %v", err)
		return
	}

	language := "go"

	fmt.Printf("Hello %v, current environemnt %v\n", language, config.Environment)

	gobasics.Test()
	gobasics.Functions()

}

type Config struct {
	Environment string `env:"ENVIRONMENT"`
}

func (config *Config) Load() error {
	return env.Parse(config)
}
