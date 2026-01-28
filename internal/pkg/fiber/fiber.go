package fiber

import (
	erros "Hog-auth/internal/app/adapters/primary/http-adapter/errors"
	"Hog-auth/internal/app/configs"
	"Hog-auth/internal/pkg/validator"
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

func NewServer(cfg *configs.Fiber, validator *validator.FiberValidator) *fiber.App {
	app := fiber.New(fiber.Config{
		StructValidator: validator,
		ErrorHandler:    erros.ErrorHandler,
		ReadTimeout:     cfg.ReadTimeout,
		WriteTimeout:    cfg.WriteTimeout,
		IdleTimeout:     cfg.IdleTimeout,
	})

	return app
}

func StartFiber(lc fx.Lifecycle, fiber *fiber.App, cfg *configs.Fiber) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := fiber.Listen(cfg.Port); err != nil {
					log.Println("dd")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return fiber.Shutdown()
		},
	})
}
