package config

import (
	"fmt"

	"github.com/protomem/socnet/pkg/env"
)

type Config struct {
	HTTP struct {
		Addr string `env:"ADDR"`
	} `envPrefix:"HTTP__"`
}

func New() (Config, error) {
	var conf Config

	err := env.Parse(&conf)
	if err != nil {
		return Config{}, fmt.Errorf("config.New: %w", err)
	}

	return conf, nil
}
