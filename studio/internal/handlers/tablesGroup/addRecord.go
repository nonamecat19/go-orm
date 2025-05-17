package tablesGroup

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	coreUtils "github.com/nonamecat19/go-orm/core/utils"
	"github.com/nonamecat19/go-orm/orm/lib/querybuilder"
	"github.com/nonamecat19/go-orm/studio/internal/utils"
	"reflect"
)

func AddTableRecord(c *fiber.Ctx) error {
	sharedData := utils.GetSharedData(c)
	tableID := c.Params("table_id")

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
		Debug().
		InsertMap(currentTable, mapFields)

	fmt.Println(err)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error adding record: %v", err))
	}

	c.Set("HX-Redirect", "/tables/"+tableID)
	return c.SendString("Record added successfully")
}
