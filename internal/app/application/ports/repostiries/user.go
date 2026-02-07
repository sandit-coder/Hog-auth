package repostiries

import (
	"Hog-auth/internal/app/domain/entities"
	"Hog-auth/internal/app/domain/vo"
	"context"

	"github.com/google/uuid"
)

type User interface {
	Create(ctx context.Context, entity *entities.User) (uuid.UUID, error)
	Update(ctx context.Context, entity *entities.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByPhoneNumber(ctx context.Context, phoneNumber vo.PhoneNumber) (*entities.User, error)
	GetByEmail(ctx context.Context, email vo.Email) (*entities.User, error)
	GetById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	Get(ctx context.Context) (*[]entities.User, error)
}
