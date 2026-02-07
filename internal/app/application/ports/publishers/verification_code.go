package publishers

import (
	"Hog-auth/internal/app/domain/events"
	"context"
)

type VerificationCode interface {
	Publish(ctx context.Context, event events.UserVerificationRequested) error
}
