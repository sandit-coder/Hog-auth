package auth

import (
	"Hog-auth/internal/app/adapters/primary/http-adapter/requests"
	"Hog-auth/internal/app/application/dto"
	"Hog-auth/internal/app/domain/types"

	"github.com/gofiber/fiber/v3"
)

func (handler *Handler) PhoneNumberVerification(c fiber.Ctx) error {
	var req requests.PhoneNumberVerification

	if err := c.Bind().Body(&req); err != nil {
		return err
	}

	verificationDto := dto.Verify{
		Code:       req.Code,
		Role:       req.Role,
		Credential: req.PhoneNumber,
		Type:       types.PhoneRegistrationType,
	}

	id, tokens, err := handler.service.Verify(&verificationDto, c.Context())
	if err != nil {
		return err
	}

	handler.cookie.SetAuthTokens(c, tokens)
	return c.Status(fiber.StatusCreated).JSON(id)
}
