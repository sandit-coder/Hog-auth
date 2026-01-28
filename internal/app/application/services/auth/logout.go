package auth

import (
	"context"
	"fmt"
)

func (auth *Auth) Logout(refreshToken string, ctx context.Context) error {
	if refreshToken == "" {
		return fmt.Errorf("refresh token is empty")
	}

	vals, err := auth.redis.HMGet(ctx, refreshToken, UserId, Active).Result()
	if err != nil {
		return fmt.Errorf("invalid refresh token")
	}

	_, ok := vals[0].(string)
	if !ok {
		return fmt.Errorf("invalid revoked field")
	}

	err = auth.redis.HSet(ctx, refreshToken, Revoked, Revoked).Err()
	if err != nil {
		return fmt.Errorf("invalid revoked field")
	}

	return nil
}
