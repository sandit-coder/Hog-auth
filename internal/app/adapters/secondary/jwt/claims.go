package jwt

import (
	"Hog-auth/internal/app/domain/vo"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserId uuid.UUID
	Role   vo.UserType
	jwt.RegisteredClaims
}

func NewClaims(customerID uuid.UUID, role vo.UserType, lyfetime time.Duration) Claims {
	now := time.Now()
	return Claims{
		UserId: customerID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(lyfetime)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "hog-auth",
			Subject:   "access_token",
		},
	}
}
