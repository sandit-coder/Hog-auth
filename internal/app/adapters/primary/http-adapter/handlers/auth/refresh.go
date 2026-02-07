package auth

import (
	"Hog-auth/internal/app/adapters/secondary/jwt"
	"time"

	"github.com/gofiber/fiber/v3"
)

func (handler *Handler) Refresh(c fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")

	accessToken, err := handler.service.Refresh(c.Context(), refreshToken)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     jwt.AcccessToken,
		Value:    accessToken,
		SameSite: Strict,
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(jwt.ExpireRefreshToken),
	})

	return c.SendStatus(fiber.StatusOK)
}
