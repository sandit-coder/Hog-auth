package provider

import (
	"Hog-auth/internal/app/application/dto"
	"Hog-auth/internal/app/application/dto/claims"
	"Hog-auth/internal/app/domain/vo"

	"github.com/google/uuid"
)

type Jwt interface {
	Parse(token string) (*claims.UserClaims, error)
	GenerateAccessToken(userId uuid.UUID, role vo.Role) (string, error)
	GenerateRefreshToken() (string, error)
	GenerateAuthTokens(userId uuid.UUID, role vo.Role) (*dto.Tokens, error)
}
