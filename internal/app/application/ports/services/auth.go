package services

import (
	"Hog-auth/internal/app/application/dto"
	"context"

	"github.com/google/uuid"
)

type Auth interface {
	Login(loginDto *dto.Login, ctx context.Context) error
	Verify(verificationDto *dto.Verify, ctx context.Context) (uuid.UUID, *dto.Tokens, error)
	Logout(refreshToken string, ctx context.Context) error
	Refresh(refreshToken string, ctx context.Context) (string, error)
}
