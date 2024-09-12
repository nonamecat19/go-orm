package app

import (
	"context"
	"github.com/a-h/templ"
	_ "github.com/a-h/templ"
	"github.com/gobuffalo/packr/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"studio/internal/config"
	"studio/internal/view/tables"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()
	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, tables.Index())
	})

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: packr.New("Assets Box", "."),
	}))

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
