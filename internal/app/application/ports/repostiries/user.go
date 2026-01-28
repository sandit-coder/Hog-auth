package repostiries

import (
	"Hog-auth/internal/app/domain/entities"
	"Hog-auth/internal/app/domain/vo"
	"context"

	"github.com/google/uuid"
)

type User interface {
	Create(entity *entities.User, ctx context.Context) (uuid.UUID, error)
	Update(entity *entities.User, ctx context.Context) error
	Delete(id uuid.UUID, ctx context.Context) error
	GetByPhoneNumber(phoneNumber vo.PhoneNumber, ctx context.Context) (*entities.User, error)
	GetByEmail(email vo.Email, ctx context.Context) (*entities.User, error)
	GetById(id uuid.UUID, ctx context.Context) (*entities.User, error)
	Get(ctx context.Context) (*[]entities.User, error)
}
