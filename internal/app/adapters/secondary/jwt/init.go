package jwt

import (
	"Hog-auth/internal/app/application/ports/services"
)

type Config struct {
	SecretKey string `koanf:"secret_key"`
}

type Jwt struct {
	cfg Config
}

func New(cfg Config) services.Jwt {
	return &Jwt{cfg: cfg}
}
