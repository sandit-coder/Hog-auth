package auth

import (
	"Hog-auth/internal/app/domain/vo"

	"github.com/google/uuid"
)

type Claims struct {
	UserId uuid.UUID
	Role   vo.UserType
}
