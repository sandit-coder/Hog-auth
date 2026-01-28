package claims

import (
	"Hog-auth/internal/app/domain/vo"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	CustomerId uuid.UUID
	Role       vo.Role
	jwt.RegisteredClaims
}

func NewUserClaims(customerID uuid.UUID, role vo.Role, lyfetime time.Duration) UserClaims {
	now := time.Now()
	return UserClaims{
		CustomerId: customerID,
		Role:       role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(lyfetime)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "hog-auth",
			Subject:   "access_token",
		},
	}
}
