package auth

import (
	"Hog-auth/internal/app/adapters/ports"
	"Hog-auth/internal/app/application/ports/services"
)

type Handler struct {
	service services.Auth
	cookie  ports.Cookie
}

func NewHandler(services services.Auth) *Handler {
	return &Handler{service: services}
}
