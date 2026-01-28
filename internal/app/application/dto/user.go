package dto

import (
	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID `json:"id" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"required, e164"`
	Email       string    `json:"email" validate:"required,email"`
}
