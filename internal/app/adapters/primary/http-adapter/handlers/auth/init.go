package auth

import (
	"Hog-auth/internal/app/application/ports/services"
)

type Handler struct {
	service services.Auth
}

func NewHandler(services services.Auth) *Handler {
	return &Handler{service: services}
}
