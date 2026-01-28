package auth

import (
	"Hog-auth/internal/app/adapters/primary/http-adapter/requests"
	"Hog-auth/internal/app/application/dto"
	"Hog-auth/internal/app/domain/types"

	"github.com/gofiber/fiber/v3"
)

func (handler *Handler) EmailVerification(c fiber.Ctx) error {
	var req requests.EmailVerification

	if err := c.Bind().Body(&req); err != nil {
		return err
	}

	verificationDto := dto.Verify{
		Credential: req.Email,
		Code:       req.Code,
		Role:       req.Role,
		Type:       types.EmailRegistrationType,
	}

	id, tokens, err := handler.service.Verify(&verificationDto, c.Context())
	if err != nil {
		return err
	}

	handler.cookie.SetAuthTokens(c, tokens)
	return c.Status(fiber.StatusOK).JSON(id)
}
