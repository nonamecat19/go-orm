package tablesGroup

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	entities2 "github.com/nonamecat19/go-orm/core/lib/entities"
	coreUtils "github.com/nonamecat19/go-orm/core/utils"
	"github.com/nonamecat19/go-orm/orm/lib/querybuilder"
	"github.com/nonamecat19/go-orm/studio/internal/model"
	"github.com/nonamecat19/go-orm/studio/internal/utils"
	tablesView "github.com/nonamecat19/go-orm/studio/internal/view/tables"
	"reflect"
)

func TableDetailPage(c *fiber.Ctx) error {
	sharedData := utils.GetSharedData(c)
	tableID := c.Params("id")

	currentTable := sharedData.TableMap[tableID]

	if currentTable == nil {
		// TODO: not found page
	}

	sortField := c.Query("sort", "id")
	sortDir := c.Query("dir", "asc")

	entityType := reflect.TypeOf(currentTable)
	sliceType := reflect.SliceOf(entityType)
	records := reflect.New(sliceType).Interface()

	_ = querybuilder.CreateQueryBuilder(sharedData.DbClient).
		OrderBy(fmt.Sprintf("%s %s", sortField, sortDir)).
		FindMany(records)

	entityFields, _ := coreUtils.GetEntityFields(reflect.New(entityType).Interface())
	systemFields := coreUtils.GetSystemFields()

	fields := make([]model.FieldInfo, len(systemFields)+len(entityFields))

	for i, fieldName := range systemFields {
		fieldNameStr, _ := coreUtils.GetFieldNameByTagValue(reflect.TypeOf(entities2.Model{}), fieldName)
		field, _ := reflect.TypeOf(entities2.Model{}).FieldByName(fieldNameStr)
		fieldType := field.Type.String()
		fields[i] = model.FieldInfo{
			Name:          fieldName,
			Type:          fieldType,
			IsSorted:      sortField == fieldName,
			SortDirection: sortDir,
		}
	}

	for i, fieldName := range entityFields {
		fieldNameStr, _ := coreUtils.GetFieldNameByTagValue(entityType, fieldName)
		field, _ := entityType.FieldByName(fieldNameStr)
		fieldType := field.Type.String()
		fields[len(systemFields)+i] = model.FieldInfo{
			Name:          fieldName,
			Type:          fieldType,
			IsSorted:      sortField == fieldName,
			SortDirection: sortDir,
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

	if c.Get("HX-Request") == "true" {
		return utils.Render(c, tablesView.TableViewContent(props.Fields, props.Data, tableID))
	}

	return utils.Render(c, tablesView.TableDetailPage(props))
}
