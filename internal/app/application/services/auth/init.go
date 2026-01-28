package auth

import (
	"Hog-auth/internal/app/application/ports/provider"
	"Hog-auth/internal/app/application/ports/repostiries"
	"Hog-auth/internal/app/application/ports/services"
	"Hog-auth/internal/app/application/services/auth/strategy_handlers"
	"Hog-auth/internal/app/domain/types"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
	"github.com/redis/go-redis/v9"
)

type Auth struct {
	trm   trm.Manager
	repo  repostiries.User
	redis *redis.Client
	jwt   provider.Jwt

	strategies map[types.RegistrationType]strategy_handlers.Strategy
}

func New(trm trm.Manager, repo repostiries.User, jwt provider.Jwt, redis *redis.Client) services.Auth {
	strategies := map[types.RegistrationType]strategy_handlers.Strategy{
		types.PhoneRegistrationType: strategy_handlers.NewPhoneStrategy(redis),
		types.EmailRegistrationType: strategy_handlers.NewEmailStrategy(redis),
	}
	return &Auth{
		trm:        trm,
		repo:       repo,
		redis:      redis,
		jwt:        jwt,
		strategies: strategies,
	}
}
