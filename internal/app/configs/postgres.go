package configs

import (
	"Hog-auth/internal/pkg/postgres"
	"errors"
	"os"
)

func NewPostgres() (*postgres.Config, error) {
	dsn := os.Getenv(DsnPostgres)
	if dsn == "" {
		return nil, errors.New("DATABASE_URL environment variable not set")
	}

	return &postgres.Config{
		Dsn: dsn,
	}, nil
}
