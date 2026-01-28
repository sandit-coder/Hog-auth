package erros

import "github.com/gofiber/fiber/v3"

func errResponse(err error) fiber.Map {
	return fiber.Map{
		"error": err.Error(),
	}
}
