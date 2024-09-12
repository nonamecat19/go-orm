package app

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"studio/internal/config"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()
	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(cfg.ServerAddr)
	if err != nil {
		return err
	}

	return nil
}
