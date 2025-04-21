package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	appUtils "github.com/nonamecat19/go-orm/core/utils"
	"github.com/nonamecat19/go-orm/studio/internal/utils"
	"github.com/nonamecat19/go-orm/studio/internal/view/settings"
	tablesView "github.com/nonamecat19/go-orm/studio/internal/view/tables"
)

func TablesPage(c *fiber.Ctx) error {
	sharedData := utils.GetSharedData(c)

	tables := make([]tablesView.Table, len(sharedData.Tables))
	for i, table := range sharedData.Tables {
		name := table.Info()
		tables[i] = tablesView.Table{
			Title: appUtils.ToHumanCase(name),
			ID:    name,
		}
	}

	props := tablesView.TablePageProps{
		Tables: tables,
	}

	return Render(c, tablesView.TablesPage(props))
}

func SettingsPage(c *fiber.Ctx) error {
	return Render(c, settings.SettingsPage())
}

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
