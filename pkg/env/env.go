package env

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

func Load(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return fmt.Errorf("env.Load: %w", err)
	}

	return nil
}

func Parse(v any) error {
	err := env.Parse(v)
	if err != nil {
		return fmt.Errorf("env.Parse: %w", err)
	}

	return nil
}
