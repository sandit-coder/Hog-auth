package routes

import (
	"Hog-auth/internal/app/adapters/primary/http-adapter/handlers/auth"

	"github.com/gofiber/fiber/v3"
)

func AppendAuth(app *fiber.App, h *auth.Handler) {
	api := app.Group("/auth/v1")

	api.Post("/login/email", h.RegisterByEmail)
	api.Post("/login/phone", h.RegisterByPhoneNumber)

	api.Post("/verification/email", h.EmailVerification)
	api.Post("/verification/phone", h.EmailVerification)

	api.Post("/refresh", h.Refresh)
	api.Post("/logout", h.Logout)
}
