package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserCredential struct {
	id             uuid.UUID
	userID         uuid.UUID
	credentialType string
	credential     string
	createdAt      time.Time
}

func (u UserCredential) ID() uuid.UUID {
	return u.id
}
func (u UserCredential) UserId() uuid.UUID {
	return u.userID
}

func (u UserCredential) CredentialType() string {
	return u.credentialType
}

func (u UserCredential) Credential() string {
	return u.credential
}

func (u UserCredential) CreatedAt() time.Time {
	return u.createdAt
}

func NewUserCredential(id uuid.UUID, userID uuid.UUID, credentialType string, credential string, createdAt time.Time) UserCredential {
	return UserCredential{
		id:             id,
		userID:         userID,
		credential:     credential,
		credentialType: credentialType,
		createdAt:      createdAt,
	}
}
