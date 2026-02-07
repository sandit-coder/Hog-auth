package jwt

import (
	"Hog-auth/internal/app/application/dto"
	"Hog-auth/internal/app/domain/auth"
	"Hog-auth/internal/app/domain/vo"
	"crypto/rand"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (j *Jwt) Parse(stringToken string) (*auth.Claims, error) {
	var claims Claims

	_, err := jwt.ParseWithClaims(stringToken, &claims, func(token *jwt.Token) (any, error) {
		return []byte(j.cfg.SecretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Alg()}))
	if err != nil {
		return nil, err
	}

	return &auth.Claims{
		UserId: claims.UserId,
		Role:   claims.Role,
	}, nil
}

func (j *Jwt) GenerateAccessToken(userId uuid.UUID, role vo.UserType) (string, error) {
	userClaims := NewClaims(userId, role, ExpireAccessToken)

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, userClaims).SignedString(j.cfg.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *Jwt) GenerateRefreshTokenString() (string, error) {
	bytes := make([]byte, 48)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

func (j *Jwt) GenerateAuthTokens(userId uuid.UUID, role vo.UserType) (*dto.Tokens, error) {
	accessToken, err := j.GenerateAccessToken(userId, role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := j.GenerateRefreshTokenString()
	if err != nil {
		return nil, err
	}

	return &dto.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
