package dto

import (
	"Hog-auth/internal/app/domain/types"
)

type Verify struct {
	Type       types.RegistrationType
	Code       string `json:"code"`
	Role       string `json:"role"`
	Credential string `json:"credentials"`
}
