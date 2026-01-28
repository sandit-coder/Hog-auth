package erros

import (
	domainErr "Hog-auth/internal/app/domain/errors"
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(c fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, domainErr.InvalidInput):
		return c.Status(fiber.StatusBadRequest).JSON(errResponse(err))

	case errors.Is(err, domainErr.AlreadyExists):
		return c.Status(fiber.StatusConflict).JSON(errResponse(err))

	case errors.Is(err, domainErr.NotFound):
		return c.Status(fiber.StatusNotFound).JSON(errResponse(err))

	case errors.Is(err, fiber.ErrUnprocessableEntity):
		return c.Status(fiber.StatusBadRequest).JSON(errResponse(errors.New("invalid request body")))

	case errors.Is(err, context.Canceled),
		errors.Is(err, context.DeadlineExceeded):
		return c.Status(fiber.StatusRequestTimeout).JSON(errResponse(errors.New("context deadline exceeded")))

	case errors.Is(err, &validator.ValidationErrors{}):
		details := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			details[e.Field()] = e.Tag()
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "validation failed",
			"field":   details,
		})

	default:
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse(err))
	}
}
