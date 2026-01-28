package auth

import (
	"Hog-auth/internal/app/adapters/secondary/providers/jwt"
	"Hog-auth/internal/app/application/dto"
	"Hog-auth/internal/app/domain/vo"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
)

func (auth *Auth) Verify(verificationDto *dto.Verify, ctx context.Context) (uuid.UUID, *dto.Tokens, error) {
	handler, ok := auth.strategies[verificationDto.Type]
	if !ok {
		return uuid.Nil, nil, fmt.Errorf("unknown verification type")
	}

	role, err := vo.NewRole(verificationDto.Role)
	if err != nil {
		return uuid.Nil, nil, err
	}

	secret, err := auth.redis.Get(ctx, verificationDto.Credential).Result()
	if err != nil {
		return uuid.Nil, nil, err
	}

	if !totp.Validate(verificationDto.Code, secret) {
		return uuid.Nil, nil, fmt.Errorf("invalid code")
	}

	err = auth.redis.Del(ctx, verificationDto.Credential).Err()
	if err != nil {
		return uuid.Nil, nil, err
	}

	user, err := handler.Verify(verificationDto.Credential, role)
	if err != nil {
		return uuid.Nil, nil, err
	}

	var createdId uuid.UUID

	err = auth.trm.Do(ctx, func(ctx context.Context) error {
		id, err := auth.repo.Create(user, ctx)
		if err != nil {
			return err
		}

		createdId = id
		return nil
	})

	tokens, err := auth.jwt.GenerateAuthTokens(createdId, role)
	if err != nil {
		return uuid.Nil, nil, err
	}

	err = auth.redis.HSet(ctx, tokens.RefreshToken, map[string]interface{}{
		UserId: user.ID().String(),
		Active: Active,
		Role:   role,
	}).Err()

	if err != nil {
		return uuid.Nil, nil, err
	}

	err = auth.redis.Expire(ctx, tokens.RefreshToken, jwt.ExpireRefreshToken).Err()
	if err != nil {
		return uuid.Nil, nil, err
	}

	return createdId, tokens, nil
}
