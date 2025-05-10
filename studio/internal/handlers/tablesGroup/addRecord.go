package tablesGroup

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	coreUtils "github.com/nonamecat19/go-orm/core/utils"
	"github.com/nonamecat19/go-orm/studio/internal/utils"
	"reflect"
)

func AddTableRecord(c *fiber.Ctx) error {
	sharedData := utils.GetSharedData(c)
	tableID := c.Params("id")

	currentTable := sharedData.TableMap[tableID]
	if currentTable == nil {
		return c.Status(fiber.StatusNotFound).SendString("Table not found")
	}

	entityType := reflect.TypeOf(currentTable)

	entityFields, _ := coreUtils.GetEntityFields(reflect.New(entityType).Interface())

	fmt.Println("Received form data for table:", tableID)
	for _, field := range entityFields {
		fieldValue := c.FormValue(field)
		fmt.Printf("Field: %s, Value: %s\n", field, fieldValue)
	}

	// TODO: Implement record creation using the appropriate querybuilder method
	c.Set("HX-Redirect", "/tables/"+tableID)
	return c.SendString("Record added successfully")
}
