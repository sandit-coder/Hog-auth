package auth

import (
	"Hog-auth/internal/app/adapters/primary/http-adapter/requests"
	"Hog-auth/internal/app/application/dto"

	"github.com/gofiber/fiber/v3"
)

func (handler *Handler) Login(c fiber.Ctx) error {
	var req requests.Login

	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}

	loginDto := dto.Login{
		Credential: req.Credential,
		Type:       req.Type,
	}

	err = handler.service.Login(c.Context(), &loginDto)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
