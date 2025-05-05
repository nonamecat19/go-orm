package handlers

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/nonamecat19/go-orm/app/entities"
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

	if _, ok := sharedData.TableMap[tableID]; !ok {
		// TODO: not found page
		//return Render(c, tablesView.TableDetailPage(props))
	}

	var orders []entities.Order
	_ = querybuilder.CreateQueryBuilder(sharedData.DbClient).FindMany(&orders)

	entityFields, _ := coreUtils.GetEntityFields(&entities.Order{})
	systemFields := coreUtils.GetSystemFields()
	fields := append(systemFields, entityFields...)

	dataSlice := make([][]string, len(orders))
	for i, record := range orders {
		var values []string

		for _, field := range systemFields {
			v := reflect.ValueOf(record)
			fieldName, _ := coreUtils.GetFieldNameByTagValue(reflect.TypeOf(entities2.Model{}), field)
			value := coreUtils.StringifyReflectValue(v.FieldByName(fieldName))
			values = append(values, fmt.Sprint(value))
		}

		for _, field := range entityFields {
			fieldName, _ := coreUtils.GetFieldNameByTagValue(reflect.TypeOf(record), field)
			value := coreUtils.StringifyReflectValue(reflect.ValueOf(record).FieldByName(fieldName))
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

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
