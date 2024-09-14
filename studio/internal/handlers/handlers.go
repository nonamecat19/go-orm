package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/nonamecat19/go-orm/studio/internal/view/settings"
	"github.com/nonamecat19/go-orm/studio/internal/view/tables"
)

func TablesPage(c *fiber.Ctx) error {
	//config.PostgresConfig

	props := tables.TablePageProps{
		Tables: []tables.Table{
			{Title: "Users", ID: "1"},
			{Title: "Orders", ID: "2"},
			{Title: "Products", ID: "3"},
		},
	}

	return Render(c, tables.TablesPage(props))
}

func SettingsPage(c *fiber.Ctx) error {
	return Render(c, settings.SettingsPage())
}

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
