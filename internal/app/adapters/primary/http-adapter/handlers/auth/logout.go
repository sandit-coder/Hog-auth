package auth

import (
	"github.com/gofiber/fiber/v3"
)

func (handler *Handler) Logout(c fiber.Ctx) error {
	refreshToken, err := handler.cookie.GetRefreshToken(c)
	if err != nil {
		return err
	}

	err = handler.service.Logout(refreshToken, c.Context())
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
