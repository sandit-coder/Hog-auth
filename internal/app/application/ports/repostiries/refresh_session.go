package repostiries

import (
	"Hog-auth/internal/app/domain/entities"
	"context"

	"github.com/google/uuid"
)

type RefreshSession interface {
	Create(ctx context.Context, entity *entities.RefreshSession) error
	Delete(ctx context.Context, id uuid.UUID) error
}
