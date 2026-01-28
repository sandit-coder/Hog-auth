package auth

import "github.com/gofiber/fiber/v3"

func (handler *Handler) Refresh(c fiber.Ctx) error {
	refreshToken, err := handler.cookie.GetRefreshToken(c)
	if err != nil {
		return err
	}

	accessToken, err := handler.service.Refresh(refreshToken, c.Context())
	if err != nil {
		return err
	}

	handler.cookie.SetAccessToken(accessToken, c)

	return c.SendStatus(fiber.StatusOK)
}
