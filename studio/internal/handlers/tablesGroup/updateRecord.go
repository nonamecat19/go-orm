package tablesGroup

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	coreUtils "github.com/nonamecat19/go-orm/core/utils"
	"github.com/nonamecat19/go-orm/orm/lib/querybuilder"
	"github.com/nonamecat19/go-orm/studio/internal/utils"
	"reflect"
)

func UpdateTableRecord(c *fiber.Ctx) error {
	sharedData := utils.GetSharedData(c)
	tableID := c.Params("id")

	currentTable := sharedData.TableMap[tableID]
	if currentTable == nil {
		return c.Status(fiber.StatusNotFound).SendString("Table not found")
	}

	entityType := reflect.TypeOf(currentTable)

	entityFields, _ := coreUtils.GetEntityFields(reflect.New(entityType).Interface())

	mapFields := make(map[string]any)

	for _, field := range entityFields {
		fieldValue := c.FormValue(field)
		if fieldValue != "" {
			mapFields[field] = fieldValue
		}
	}

	err := querybuilder.CreateQueryBuilder(sharedData.DbClient).
		Where("id = ?", c.Params("id")).
		UpdateMap(currentTable, mapFields)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error updating record: %v", err))
	}

	c.Set("HX-Redirect", "/tables/"+tableID)
	return c.SendString("Record updated successfully")
}
