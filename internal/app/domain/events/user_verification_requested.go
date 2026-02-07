package events

import (
	"time"

	"github.com/google/uuid"
)

type UserVerificationRequested struct {
	EventId          uuid.UUID
	Credential       string
	VerificationType string
	Code             string
	OccurredAt       time.Time
}
