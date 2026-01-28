package ports

import (
	"Hog-auth/internal/app/application/dto"

	"github.com/gofiber/fiber/v3"
)

type Cookie interface {
	SetAuthTokens(c fiber.Ctx, tokens *dto.Tokens)
	SetAccessToken(accessToken string, c fiber.Ctx)
	ClearAuthTokens(c fiber.Ctx)
	GetRefreshToken(c fiber.Ctx) (string, error)
	GetAccessToken(c fiber.Ctx) (string, error)
}
