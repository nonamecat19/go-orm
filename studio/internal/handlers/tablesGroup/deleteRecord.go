package tablesGroup

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	coreEntities "github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/orm/lib/querybuilder"
	"github.com/nonamecat19/go-orm/studio/internal/utils"
	"reflect"
)

func GetEntityPointer(entity coreEntities.IEntity) interface{} {
	entityType := reflect.TypeOf(entity).Elem()
	entityPtr := reflect.New(entityType)
	return entityPtr
}

func DeleteTableRecord(c *fiber.Ctx) error {
	sharedData := utils.GetSharedData(c)
	tableID := c.Params("table_id")

	currentTable := sharedData.TableMap[tableID]
	if currentTable == nil {
		return c.Status(fiber.StatusNotFound).SendString("Table not found")
	}

	entityType := reflect.TypeOf(currentTable)
	records := reflect.New(entityType).Interface()

	err := querybuilder.CreateQueryBuilder(sharedData.DbClient).
		Where("id = ?", c.Params("record_id")).
		DeleteMany(records)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error updating record: %v", err))
	}

	c.Set("HX-Redirect", sharedData.Prefix+"/tables/"+tableID)
	return c.SendString("Record deleted successfully")
}
