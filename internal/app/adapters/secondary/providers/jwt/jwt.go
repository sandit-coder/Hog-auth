package jwt

import (
	"Hog-auth/internal/app/application/dto"
	"Hog-auth/internal/app/application/dto/claims"
	"Hog-auth/internal/app/domain/vo"
	"crypto/rand"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (j *Provider) Parse(stringToken string) (*claims.UserClaims, error) {
	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (any, error) {
		return []byte(j.SecretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Alg()}))
	if err != nil {
		return nil, err
	}

	return token.Claims.(*claims.UserClaims), nil
}

func (j *Provider) GenerateAccessToken(userId uuid.UUID, role vo.Role) (string, error) {
	userClaims := claims.NewUserClaims(userId, role, ExpireAccessToken)

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, userClaims).SignedString(j.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *Provider) GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 48)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

func (j *Provider) GenerateAuthTokens(userId uuid.UUID, role vo.Role) (*dto.Tokens, error) {
	accessToken, err := j.GenerateAccessToken(userId, role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := j.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &dto.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
