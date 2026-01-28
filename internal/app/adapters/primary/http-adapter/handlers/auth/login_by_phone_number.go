package auth

import (
	"Hog-auth/internal/app/adapters/primary/http-adapter/requests"
	"Hog-auth/internal/app/application/dto"
	"Hog-auth/internal/app/domain/types"

	"github.com/gofiber/fiber/v3"
)

func (handler *Handler) RegisterByPhoneNumber(c fiber.Ctx) error {
	var req requests.RegisterByPhoneNumber

	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}

	registrationDto := dto.Login{
		Credential: req.PhoneNumber,
		Type:       types.PhoneRegistrationType,
	}

	err = handler.service.Login(&registrationDto, c.Context())
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
