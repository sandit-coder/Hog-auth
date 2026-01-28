package jwt

import (
	"errors"
	"os"
)

func GetSecretKey() (string, error) {
	secretKey := os.Getenv(SecretKey)

	if secretKey == "" {
		return "", errors.New("missing secret key")
	}
	return secretKey, nil
}
