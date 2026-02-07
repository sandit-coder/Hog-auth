package auth

import (
	"Hog-auth/internal/app/application/dto"
	"Hog-auth/internal/app/domain/entities"
	"Hog-auth/internal/app/domain/types"
	"Hog-auth/internal/app/domain/vo"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
)

func (auth *Auth) Verify(ctx context.Context, verificationDto *dto.Verify) (uuid.UUID, *dto.Tokens, error) {
	regType, err := types.CredentialTypeFromString(verificationDto.Type)
	if err != nil {
		return uuid.Nil, nil, err
	}

	normalizer, ok := auth.strategies[regType]
	if !ok {
		return uuid.Nil, nil, fmt.Errorf("unknown verification type")
	}

	userType, err := vo.NewUserType(verificationDto.Role)
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

	normalizeCredential, err := normalizer.NormalizeCredential(verificationDto.Credential)
	if err != nil {
		return uuid.Nil, nil, err
	}

	userID := uuid.New()
	credential := entities.NewUserCredential(uuid.New(), userID, strings.ToLower(verificationDto.Type), normalizeCredential, time.Now())

	user, err := entities.NewUser(uuid.New(), userType, credential)
	if err != nil {
		return uuid.Nil, nil, err
	}

	err = auth.tx.Do(ctx, func(ctx context.Context) error {
		_, err := auth.userRepo.Create(ctx, user)
		if err != nil {
			return err
		}

		return nil
	})

	err = auth.redis.Del(ctx, verificationDto.Credential).Err()
	if err != nil {
		return uuid.Nil, nil, err
	}

	tokens, err := auth.jwt.GenerateAuthTokens(userID, userType)
	if err != nil {
		return uuid.Nil, nil, err
	}

	refreshTokenHash, err := vo.NewRefreshTokenHash(tokens.RefreshToken)
	if err != nil {
		return uuid.Nil, nil, err
	}

	refreshSession, err := entities.NewRefreshSession(uuid.New(), userID, refreshTokenHash, userType)
	if err != nil {
		return uuid.Nil, nil, err
	}

	err = auth.tx.Do(ctx, func(ctx context.Context) error {
		err := auth.refreshSessionRepo.Create(ctx, refreshSession)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return uuid.Nil, nil, err
	}

	return userID, tokens, nil
}
