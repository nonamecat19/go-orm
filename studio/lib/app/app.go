package app

import (
	_ "github.com/a-h/templ"
	"github.com/gobuffalo/packr/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	coreEntities "github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/orm/lib/client"
	"github.com/nonamecat19/go-orm/studio/internal/handlers/settingsGroup"
	"github.com/nonamecat19/go-orm/studio/internal/handlers/tablesGroup"
	"github.com/nonamecat19/go-orm/studio/internal/utils"
	"strings"
)

func AddStudioFiberGroup(app *fiber.App, tables []coreEntities.IEntity, client client.DbClient, prefix string) {
	studioGroup := app.Group(prefix)

	sharedData := fiber.Map{
		"tableMap":     utils.GetTableMap(tables),
		"tables":       tables,
		"client":       client,
		"prefix":       prefix,
		"assetsPrefix": prefix + "/assets/web/public/assets",
	}

	studioGroup.Use(func(c *fiber.Ctx) error {
		c.Locals("data", sharedData)
		err := c.Next()

		body := c.Response().Body()
		bodyStr := string(body)

		bodyStr = strings.Replace(bodyStr, "action=\"/api/tables/", "action=\""+prefix+"/api/tables/", -1)
		bodyStr = strings.Replace(bodyStr, "hx-get=\"/tables/", "hx-get=\""+prefix+"/tables/", -1)
		bodyStr = strings.Replace(bodyStr, "hx-post=\"/api/tables/", "hx-post=\""+prefix+"/api/tables/", -1)
		bodyStr = strings.Replace(bodyStr, "hx-put=\"/api/tables/", "hx-put=\""+prefix+"/api/tables/", -1)
		bodyStr = strings.Replace(bodyStr, "hx-delete=\"/api/tables/", "hx-delete=\""+prefix+"/api/tables/", -1)
		bodyStr = strings.Replace(bodyStr, "href=\"/", "href=\""+prefix+"/", -1)

		c.Response().SetBody([]byte(bodyStr))
		return err
	})

	studioGroup.Get("/", tablesGroup.TablesPage)
	studioGroup.Get("/tables/:id", tablesGroup.TableDetailPage)
	studioGroup.Get("/settings", settingsGroup.SettingsPage)

	studioGroup.Post("/api/tables/:table_id/records", tablesGroup.AddTableRecord)
	studioGroup.Put("/api/tables/:table_id/records/:record_id", tablesGroup.UpdateTableRecord)
	studioGroup.Delete("/api/tables/:table_id/records/:record_id", tablesGroup.DeleteTableRecord)

	studioGroup.Use("/assets", filesystem.New(filesystem.Config{
		Root: packr.New("Assets Box", "."),
	}))
}
