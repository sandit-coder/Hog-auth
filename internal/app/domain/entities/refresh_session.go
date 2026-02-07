package entities

import (
	"Hog-auth/internal/app/domain/vo"

	"github.com/google/uuid"
)

type RefreshSession struct {
	id               uuid.UUID
	userId           uuid.UUID
	refreshTokenHash vo.RefreshTokenHash
	userType         vo.UserType
	revoked          bool
}

func (r *RefreshSession) ID() uuid.UUID {
	return r.id
}

func (r *RefreshSession) UserId() uuid.UUID {
	return r.userId
}

func (r *RefreshSession) RefreshTokenHash() vo.RefreshTokenHash {
	return r.refreshTokenHash
}

func (r *RefreshSession) UserType() vo.UserType {
	return r.userType
}

func (r *RefreshSession) Revoked() bool {
	return r.revoked
}

func NewRefreshSession(
	id uuid.UUID,
	userId uuid.UUID,
	refreshTokenHash vo.RefreshTokenHash,
	role vo.UserType) (*RefreshSession, error) {

	return &RefreshSession{
		id:               id,
		userId:           userId,
		refreshTokenHash: refreshTokenHash,
		userType:         role,
		revoked:          false,
	}, nil
}

func (r *RefreshSession) Cancel() {
	r.revoked = true
}
