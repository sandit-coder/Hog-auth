package cookie

import (
	"Hog-auth/internal/app/adapters/secondary/providers/jwt"
	"Hog-auth/internal/app/application/dto"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

const (
	Strict = "Strict"
)

func (provider *Provider) SetAuthTokens(dto *dto.Tokens) {
	provider.c.Cookie(&fiber.Cookie{
		Name:     jwt.AcccessToken,
		Value:    dto.AccessToken,
		SameSite: Strict,
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(jwt.ExpireRefreshToken),
	})

	provider.c.Cookie(&fiber.Cookie{
		Name:     jwt.RefreshToken,
		Value:    dto.RefreshToken,
		SameSite: Strict,
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(jwt.ExpireAccessToken),
	})
}

func ClearAuthTokens(c fiber.Ctx) {
	c.ClearCookie(jwt.AcccessToken)
	c.ClearCookie(jwt.RefreshToken)
}

func GetRefreshToken(c fiber.Ctx) (string, error) {
	refreshToken := c.Cookies(jwt.RefreshToken)
	if refreshToken == "" {
		return "", fmt.Errorf("in cookie refresh token is empty")
	}

	return refreshToken, nil
}

func SetAccessToken(accessToken string, c fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     jwt.AcccessToken,
		Value:    accessToken,
		SameSite: Strict,
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(jwt.ExpireRefreshToken),
	})
}
