package strategy_handlers

import (
	"Hog-auth/internal/app/domain/entities"
	"Hog-auth/internal/app/domain/vo"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/redis/go-redis/v9"
)

const ()

type EmailStrategy struct {
	redis *redis.Client
}

func NewEmailStrategy(redis *redis.Client) *EmailStrategy {
	return &EmailStrategy{
		redis: redis}
}

func (s *EmailStrategy) Initiate(email string, ctx context.Context) error {
	voEmail, err := vo.NewEmail(email)
	if err != nil {
		return err
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "hog-auth",
		AccountName: voEmail.Value,
	})
	if err != nil {
		return err
	}

	code, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return err
	}

	return s.redis.Set(ctx, voEmail.Value, key.Secret(), time.Minute*5).Err()
}

func (s *EmailStrategy) Verify(credential string, role vo.Role) (*entities.User, error) {
	email, err := vo.NewEmail(credential)
	if err != nil {
		return nil, err
	}

	user, err := entities.NewUser(uuid.New(), role, entities.WithEmail(email))
	if err != nil {
		return nil, err
	}

	return user, nil
}
