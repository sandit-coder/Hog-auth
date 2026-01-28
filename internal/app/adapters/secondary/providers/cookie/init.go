package cookie

import "github.com/gofiber/fiber/v3"

type Provider struct {
	c fiber.Ctx
}

func New(c fiber.Ctx) *Provider {
	return &Provider{c: c}
}
