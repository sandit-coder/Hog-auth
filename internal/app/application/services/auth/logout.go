package auth

import (
	"context"
	"fmt"
)

func (auth *Auth) Logout(ctx context.Context, refreshToken string) error {
	if refreshToken == "" {
		return fmt.Errorf("session session is empty")
	}

	vals, err := auth.redis.HMGet(ctx, refreshToken, UserId, Active).Result()
	if err != nil {
		return fmt.Errorf("invalid session session")
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
