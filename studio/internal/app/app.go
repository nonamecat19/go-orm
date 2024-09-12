package app

import (
	"github.com/a-h/templ"
	_ "github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"studio/internal/config"
	"studio/internal/view/tables"
)

func Run() error {
	cfg := config.NewConfig()
	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, tables.Index())
	})

	err := app.Listen(cfg.ServerAddr)
	if err != nil {
		return err
	}

	return nil
}

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
