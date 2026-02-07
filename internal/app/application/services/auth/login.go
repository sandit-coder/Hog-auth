package auth

import (
	"Hog-auth/internal/app/application/dto"
	"Hog-auth/internal/app/domain/events"
	"Hog-auth/internal/app/domain/types"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
)

func (auth *Auth) Login(ctx context.Context, loginDto *dto.Login) error {

	credType, err := types.CredentialTypeFromString(loginDto.Type)

	if err != nil {
		return err
	}

	normalizer, ok := auth.strategies[credType]
	if !ok {
		return fmt.Errorf("unknown registration type type")
	}

	credential, err := normalizer.NormalizeCredential(loginDto.Credential)
	if err != nil {
		return err
	}

	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "hog-auth",
		AccountName: credential,
	})
	if err != nil {
		return err
	}

	code, err := totp.GenerateCode(secret.Secret(), time.Now())
	if err != nil {
		return err
	}

	err = auth.redis.Set(ctx, credential, secret, VerificationCodeLyfetime).Err()
	if err != nil {
		return err
	}

	event := events.UserVerificationRequested{
		EventId:          uuid.New(),
		VerificationType: strings.ToLower(loginDto.Type),
		Credential:       credential,
		Code:             code,
		OccurredAt:       time.Now(),
	}

	auth.publisher.Publish(ctx, event)

	return nil
}
