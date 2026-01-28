package responses

import "github.com/google/uuid"

type RegisterResponse struct {
	UserId uuid.UUID `json:"user_id" validate:"required,uuid"`
}
