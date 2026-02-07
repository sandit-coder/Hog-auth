package vo

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type RefreshTokenHash struct {
	Value []byte
}

func NewRefreshTokenHash(value string) (RefreshTokenHash, error) {
	if value == "" {
		return RefreshTokenHash{}, errors.New("value is empty")
	}

	refreshTokenHash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return RefreshTokenHash{}, err
	}

	return RefreshTokenHash{
		Value: refreshTokenHash,
	}, nil
}
