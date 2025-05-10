package handlers

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	entities2 "github.com/nonamecat19/go-orm/core/lib/entities"
	coreUtils "github.com/nonamecat19/go-orm/core/utils"
	"github.com/nonamecat19/go-orm/orm/lib/querybuilder"
	"github.com/nonamecat19/go-orm/studio/internal/utils"
	"github.com/nonamecat19/go-orm/studio/internal/view/settings"
	tablesView "github.com/nonamecat19/go-orm/studio/internal/view/tables"
	"reflect"
)

func TablesPage(c *fiber.Ctx) error {
	sharedData := utils.GetSharedData(c)

	tables := make([]tablesView.Table, len(sharedData.Tables))
	for i, table := range sharedData.Tables {
		name := table.Info()
		tables[i] = tablesView.Table{
			Title: coreUtils.ToHumanCase(name),
			ID:    name,
		}
	}

	props := tablesView.TablePageProps{
		Tables: tables,
	}

	return Render(c, tablesView.TablesPage(props))
}

func TableDetailPage(c *fiber.Ctx) error {
	sharedData := utils.GetSharedData(c)
	tableID := c.Params("id")

	currentTable := sharedData.TableMap[tableID]

	if currentTable == nil {
		// TODO: not found page
		//return Render(c, tablesView.TableDetailPage(props))
	}

	entityType := reflect.TypeOf(currentTable)
	sliceType := reflect.SliceOf(entityType)
	records := reflect.New(sliceType).Interface()
	_ = querybuilder.CreateQueryBuilder(sharedData.DbClient).FindMany(records)

	entityFields, _ := coreUtils.GetEntityFields(reflect.New(entityType).Interface())
	systemFields := coreUtils.GetSystemFields()

	fields := make([]tablesView.FieldInfo, len(systemFields)+len(entityFields))

	for i, fieldName := range systemFields {
		fieldNameStr, _ := coreUtils.GetFieldNameByTagValue(reflect.TypeOf(entities2.Model{}), fieldName)
		field, _ := reflect.TypeOf(entities2.Model{}).FieldByName(fieldNameStr)
		fieldType := field.Type.String()
		fields[i] = tablesView.FieldInfo{
			Name: fieldName,
			Type: fieldType,
		}
	}

	for i, fieldName := range entityFields {
		fieldNameStr, _ := coreUtils.GetFieldNameByTagValue(entityType, fieldName)
		field, _ := entityType.FieldByName(fieldNameStr)
		fieldType := field.Type.String()
		fields[len(systemFields)+i] = tablesView.FieldInfo{
			Name: fieldName,
			Type: fieldType,
		}
	}

	dataSlice := make([][]string, reflect.ValueOf(records).Elem().Len())
	for i := 0; i < reflect.ValueOf(records).Elem().Len(); i++ {
		var values []string
		record := reflect.ValueOf(records).Elem().Index(i)

		for _, field := range systemFields {
			fieldName, _ := coreUtils.GetFieldNameByTagValue(reflect.TypeOf(entities2.Model{}), field)
			value := coreUtils.StringifyReflectValue(record.FieldByName(fieldName))
			values = append(values, fmt.Sprint(value))
		}

		for _, field := range entityFields {
			fieldName, _ := coreUtils.GetFieldNameByTagValue(record.Type(), field)
			value := coreUtils.StringifyReflectValue(record.FieldByName(fieldName))
			values = append(values, fmt.Sprint(value))
		}

		dataSlice[i] = values
	}

	props := tablesView.TableDetailProps{
		Table: tablesView.Table{
			Title: coreUtils.ToHumanCase(tableID),
			ID:    tableID,
		},
		Data:   dataSlice,
		Fields: fields,
	}
	return Render(c, tablesView.TableDetailPage(props))
}

func SettingsPage(c *fiber.Ctx) error {
	return Render(c, settings.SettingsPage())
}

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

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
