package dto

import (
	"Hog-auth/internal/app/domain/types"
)

type Login struct {
	Type       types.RegistrationType
	Credential string
}
