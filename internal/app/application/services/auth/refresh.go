package auth

import (
	"Hog-auth/internal/app/domain/vo"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (auth *Auth) Refresh(ctx context.Context, refreshToken string) (string, error) {

	exists, err := auth.redis.Exists(ctx, refreshToken).Result()
	if err != nil {
		return "", fmt.Errorf("redis error: %w", err)
	}
	if exists == 0 {
		return "", fmt.Errorf("session not found")
	}

	vals, err := auth.redis.HMGet(ctx, refreshToken, UserId, Active, Role).Result()
	if err != nil {
		return "", err
	}

	revokedStr, ok := vals[0].(string)
	if !ok {
		return "", fmt.Errorf("invalid revoked field")
	}
	if revokedStr == Revoked {
		return "", fmt.Errorf("session revoked")
	}

	userID, err := uuid.Parse(vals[1].(string))
	if err != nil {
		return "", fmt.Errorf("invalid user_id")
	}

	role, ok := vals[2].(string)
	if !ok {
		return "", fmt.Errorf("invalid role")
	}

	voRole, err := vo.NewUserType(role)
	if err != nil {
		return "", fmt.Errorf("invalid role")
	}

	accessToken, err := auth.jwt.GenerateAccessToken(userID, voRole)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
