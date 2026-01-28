package jwt

import (
	"Hog-auth/internal/app/application/ports/provider"
	"fmt"
	"os"
)

type Provider struct {
	SecretKey string
}

func New() (provider.Jwt, error) {
	secret := os.Getenv(SecretKey)
	if secret == "" {
		return nil, fmt.Errorf("no secret key provided")
	}

	return &Provider{
		SecretKey: secret,
	}, nil
}
