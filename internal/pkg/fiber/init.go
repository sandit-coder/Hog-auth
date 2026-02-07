package fiber

import (
	erros "Hog-auth/internal/app/adapters/primary/http-adapter/errors"
	"Hog-auth/internal/pkg/validator"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

type Config struct {
	Port         string
	ReadTimeout  time.Duration `koanf:"read_timeout"`
	WriteTimeout time.Duration `koanf:"write_timeout"`
	IdleTimeout  time.Duration `koanf:"idle_timeout"`
}

func NewServer(cfg Config, validator *validator.FiberValidator) *fiber.App {
	app := fiber.New(fiber.Config{
		StructValidator: validator,
		ErrorHandler:    erros.ErrorHandler,
		ReadTimeout:     cfg.ReadTimeout,
		WriteTimeout:    cfg.WriteTimeout,
		IdleTimeout:     cfg.IdleTimeout,
	})

	return app
}

func StartFiber(lc fx.Lifecycle, fiber *fiber.App, cfg Config) {
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
