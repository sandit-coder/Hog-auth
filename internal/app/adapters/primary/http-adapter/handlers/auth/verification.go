package auth

import (
	"Hog-auth/internal/app/adapters/primary/http-adapter/requests"
	"Hog-auth/internal/app/adapters/secondary/jwt"
	"Hog-auth/internal/app/application/dto"
	"time"

	"github.com/gofiber/fiber/v3"
)

func (handler *Handler) Verification(c fiber.Ctx) error {
	var req requests.Verification

	if err := c.Bind().Body(&req); err != nil {
		return err
	}

	verificationDto := dto.Verify{
		Credential: req.Credential,
		Code:       req.Code,
		Role:       req.Role,
		Type:       req.Type,
	}

	id, tokens, err := handler.service.Verify(c.Context(), &verificationDto)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     jwt.AcccessToken,
		Value:    tokens.AccessToken,
		SameSite: Strict,
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(jwt.ExpireRefreshToken),
	})

	return c.Status(fiber.StatusOK).JSON(id)
}
