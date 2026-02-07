package services

import (
	"Hog-auth/internal/app/application/dto"

	"context"

	"github.com/google/uuid"
)

type Auth interface {
	Login(ctx context.Context, loginDto *dto.Login) error
	Verify(ctx context.Context, verificationDto *dto.Verify) (uuid.UUID, *dto.Tokens, error)
	Logout(ctx context.Context, refreshToken string) error
	Refresh(ctx context.Context, refreshToken string) (string, error)
}
