package internal

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	LogLevel string `env:"LOG_LEVEL"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("Error while parsing env config: %v", err)
	}
	return cfg, nil
}
