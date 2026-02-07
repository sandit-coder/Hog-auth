package routes

import (
	"Hog-auth/internal/app/adapters/primary/http-adapter/handlers/auth"

	"github.com/gofiber/fiber/v3"
)

func AppendAuth(app *fiber.App, h *auth.Handler) {
	api := app.Group("/jwt/v1")

	api.Post("/login", h.Login)

	api.Post("/verification", h.Verification)

	api.Post("/refresh", h.Refresh)
	api.Post("/logout", h.Logout)
}
