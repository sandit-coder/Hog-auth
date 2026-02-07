package entities

import (
	"Hog-auth/internal/app/domain/vo"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	id                    uuid.UUID
	refreshTokenSessionId uuid.UUID
	userType              vo.UserType
	credentials           []UserCredential
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) UserType() vo.UserType {
	return u.userType
}

func (u *User) RefreshTokenSessionId() uuid.UUID {
	return u.refreshTokenSessionId
}

func (u *User) Credentials() []UserCredential {
	return append([]UserCredential{}, u.credentials...)
}

func NewUser(id uuid.UUID, role vo.UserType, credentials UserCredential) (*User, error) {
	if id == uuid.Nil {
		id = uuid.New()
	}

	u := &User{
		id:       id,
		userType: role,
	}

	if err := u.AddCredential(credentials); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) AddCredential(cred UserCredential) error {
	if cred.UserId() != u.id {
		return fmt.Errorf("credential belongs to another user")
	}

	for _, c := range u.Credentials() {
		if c.Credential() == cred.Credential() {
			return fmt.Errorf("credential type already exists")
		}
	}

	u.credentials = append(u.credentials, cred)
	return nil
}
