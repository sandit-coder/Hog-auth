package auth

import (
	"Hog-auth/internal/app/adapters/primary/http-adapter/requests"
	"Hog-auth/internal/app/application/dto"
	"Hog-auth/internal/app/domain/types"

	"github.com/gofiber/fiber/v3"
)

func (handler *Handler) RegisterByEmail(c fiber.Ctx) error {
	var req requests.RegisterByEmail

	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}

	registerDto := dto.Login{
		Credential: req.Email,
		Type:       types.EmailRegistrationType,
	}

	err = handler.service.Login(&registerDto, c.Context())
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
