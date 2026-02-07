package auth

import (
	"Hog-auth/internal/app/application/ports/publishers"
	"Hog-auth/internal/app/application/ports/repostiries"
	"Hog-auth/internal/app/application/ports/services"
	"Hog-auth/internal/app/application/services/auth/strategies"
	"Hog-auth/internal/app/domain/types"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
	"github.com/redis/go-redis/v9"
)

type Auth struct {
	tx                 trm.Manager
	userRepo           repostiries.User
	refreshSessionRepo repostiries.RefreshSession
	redis              *redis.Client
	jwt                services.Jwt
	publisher          publishers.VerificationCode

	strategies map[types.CredentialType]strategies.Strategy
}

func New(
	trm trm.Manager,
	userRepo repostiries.User,
	refreshSessionRepo repostiries.RefreshSession,
	jwt services.Jwt,
	redis *redis.Client,
	publisher publishers.VerificationCode) services.Auth {

	strategies := map[types.CredentialType]strategies.Strategy{
		types.PhoneCredentialType: strategies.NewPhoneNumberNormalizer(),
		types.EmailCredentialType: strategies.NewEmailNormalizer(),
	}

	return &Auth{
		tx:                 trm,
		userRepo:           userRepo,
		refreshSessionRepo: refreshSessionRepo,
		redis:              redis,
		jwt:                jwt,
		publisher:          publisher,
		strategies:         strategies,
	}
}
