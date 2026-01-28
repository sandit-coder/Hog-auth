package configs

import (
	"Hog-auth/internal/pkg/redis"
	"fmt"
	"os"
)

func NewRedis() (*redis.Config, error) {
	dsn := os.Getenv(DsnRedis)
	if dsn == "" {
		return nil, fmt.Errorf("env variable %s is not set", dsn)
	}

	return &redis.Config{
		Dsn: dsn,
	}, nil
}
