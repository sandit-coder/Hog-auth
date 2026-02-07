package services

import (
	"Hog-auth/internal/app/application/dto"
	"Hog-auth/internal/app/domain/auth"
	"Hog-auth/internal/app/domain/vo"

	"github.com/google/uuid"
)

type Jwt interface {
	GenerateAccessToken(userId uuid.UUID, role vo.UserType) (string, error)
	GenerateRefreshTokenString() (string, error)
	GenerateAuthTokens(userId uuid.UUID, role vo.UserType) (*dto.Tokens, error)
	Parse(stringToken string) (*auth.Claims, error)
}
