package auth

import (
	"github.com/gofiber/fiber/v3"
)

func (handler *Handler) Logout(c fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")

	if err := handler.service.Logout(c.Context(), refreshToken); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
