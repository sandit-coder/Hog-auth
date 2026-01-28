package configs

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	if err := godotenv.Load(EnvFile); err != nil {
		return fmt.Errorf("errors loading .env file: %w", err)
	}
	return nil
}
