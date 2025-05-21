package utils

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	coreEntities "github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/orm/lib/client"
)

func GetTableMap(tables []coreEntities.IEntity) map[string]coreEntities.IEntity {
	tableMap := make(map[string]coreEntities.IEntity)
	for _, table := range tables {
		tableMap[table.Info()] = table
	}
	return tableMap
}

type SharedData struct {
	TableMap     map[string]coreEntities.IEntity
	Tables       []coreEntities.IEntity
	DbClient     client.DbClient
	Prefix       string
	AssetsPrefix string
}

func GetSharedData(c *fiber.Ctx) SharedData {
	data := c.Locals("data").(fiber.Map)

	tableMap := data["tableMap"].(map[string]coreEntities.IEntity)
	tables := data["tables"].([]coreEntities.IEntity)
	dbClient := data["client"].(client.DbClient)
	prefix := data["prefix"].(string)
	assetsPrefix := data["assetsPrefix"].(string)

	return SharedData{
		TableMap:     tableMap,
		Tables:       tables,
		DbClient:     dbClient,
		Prefix:       prefix,
		AssetsPrefix: assetsPrefix,
	}
}

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
