package entities

import (
	"Hog-auth/internal/app/domain/vo"

	"github.com/google/uuid"
)

type User struct {
	id          uuid.UUID
	role        vo.Role
	email       *vo.Email
	phoneNumber *vo.PhoneNumber
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Role() vo.Role {
	return u.role
}

func (u User) Email() *vo.Email {
	return u.email
}

func (u User) PhoneNumber() *vo.PhoneNumber {
	return u.phoneNumber
}

type UserOption func(*User)

func WithEmail(email vo.Email) UserOption {
	return func(u *User) {
		u.email = &email
	}
}

func WithPhoneNumber(phoneNumber vo.PhoneNumber) UserOption {
	return func(u *User) {
		u.phoneNumber = &phoneNumber
	}
}

func NewUser(id uuid.UUID, role vo.Role, opts ...UserOption) (*User, error) {
	if id == uuid.Nil {
		id = uuid.New()
	}

	u := &User{
		id:   id,
		role: role,
	}

	for _, opt := range opts {
		opt(u)
	}

	return u, nil
}
