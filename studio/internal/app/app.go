package app

import (
	"context"
	_ "github.com/a-h/templ"
	"github.com/gobuffalo/packr/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	coreEntities "github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/orm/lib/client"
	"github.com/nonamecat19/go-orm/studio/internal/config"
	"github.com/nonamecat19/go-orm/studio/internal/handlers/settingsGroup"
	"github.com/nonamecat19/go-orm/studio/internal/handlers/tablesGroup"
	"github.com/nonamecat19/go-orm/studio/internal/utils"
)

func Run(ctx context.Context, tables []coreEntities.IEntity, client client.DbClient, cfg config.StudioConfig) error {
	app := fiber.New(fiber.Config{})

	sharedData := fiber.Map{
		"tableMap": utils.GetTableMap(tables),
		"tables":   tables,
		"client":   client,
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("data", sharedData)
		return c.Next()
	})

	app.Get("/", tablesGroup.TablesPage)
	app.Get("/tables/:id", tablesGroup.TableDetailPage)
	app.Get("/settings", settingsGroup.SettingsPage)

	app.Post("/api/tables/:id/records", tablesGroup.AddTableRecord)

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: packr.New("Assets Box", "."),
	}))

	err := app.Listen(cfg.ServerAddr)
	if err != nil {
		return err
	}

	return nil
}
