package strategy_handlers

import (
	"Hog-auth/internal/app/domain/entities"
	"Hog-auth/internal/app/domain/vo"
	"context"
)

type Strategy interface {
	Initiate(credential string, ctx context.Context) error

	Verify(credential string, role vo.Role) (*entities.User, error)
}
