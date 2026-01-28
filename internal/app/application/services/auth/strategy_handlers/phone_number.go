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

type PhoneStrategy struct {
	redis *redis.Client
}

func NewPhoneStrategy(redis *redis.Client) *PhoneStrategy {
	return &PhoneStrategy{
		redis: redis}
}

func (s *PhoneStrategy) Initiate(phoneNumber string, ctx context.Context) error {
	voPhone, err := vo.NewPhoneNumber(phoneNumber)
	if err != nil {
		return err
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "hog-auth",
		AccountName: voPhone.Value,
	})
	if err != nil {
		return err
	}

	code, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return err
	}

	//в кафочку код кидаем

	return s.redis.Set(ctx, voPhone.Value, key.Secret(), time.Minute*5).Err()
}

func (s *PhoneStrategy) Verify(credential string, role vo.Role) (*entities.User, error) {
	phoneNumber, err := vo.NewPhoneNumber(credential)
	if err != nil {
		return nil, err
	}

	user, err := entities.NewUser(uuid.New(), role, entities.WithPhoneNumber(phoneNumber))
	if err != nil {
		return nil, err
	}

	return user, nil
}
