package app

import (
	"context"
	_ "github.com/a-h/templ"
	"github.com/gobuffalo/packr/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/nonamecat19/go-orm/studio/internal/config"
	"github.com/nonamecat19/go-orm/studio/internal/handlers"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()
	app := fiber.New(fiber.Config{})

	app.Get("/", handlers.TablesPage)
	app.Get("/settings", handlers.SettingsPage)

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: packr.New("Assets Box", "."),
	}))

	err := app.Listen(cfg.ServerAddr)
	if err != nil {
		return err
	}

	return nil
}
