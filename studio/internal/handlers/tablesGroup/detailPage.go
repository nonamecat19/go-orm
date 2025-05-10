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
	"strings"
)

func isFieldNullable(fieldType string) bool {
	return strings.Contains(fieldType, "*")
}

func getTableRecords(sharedData utils.SharedData, tableID, sortField, sortDir string) (interface{}, reflect.Type) {
	currentTable := sharedData.TableMap[tableID]

	if currentTable == nil {
		return nil, nil
	}

	entityType := reflect.TypeOf(currentTable)
	sliceType := reflect.SliceOf(entityType)
	records := reflect.New(sliceType).Interface()

	_ = querybuilder.CreateQueryBuilder(sharedData.DbClient).
		OrderBy(fmt.Sprintf("%s %s", sortField, sortDir)).
		FindMany(records)

	return records, entityType
}

func buildFieldInfo(entityType reflect.Type, systemFields, entityFields []string, sortField, sortDir string) []model.FieldInfo {
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
			IsNullable:    isFieldNullable(fieldType),
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
			IsNullable:    isFieldNullable(fieldType),
		}
	}

	return fields
}

func buildDataSlice(records interface{}, systemFields, entityFields []string) [][]string {
	recordsValue := reflect.ValueOf(records).Elem()
	dataSlice := make([][]string, recordsValue.Len())

	for i := 0; i < recordsValue.Len(); i++ {
		var values []string
		record := recordsValue.Index(i)

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

	return dataSlice
}

func TableDetailPage(c *fiber.Ctx) error {
	sharedData := utils.GetSharedData(c)
	tableID := c.Params("id")

	sortField := c.Query("sort", "id")
	sortDir := c.Query("dir", "asc")

	records, entityType := getTableRecords(sharedData, tableID, sortField, sortDir)
	if entityType == nil {
		// TODO: not found page
		return c.Status(fiber.StatusNotFound).SendString("Table not found")
	}

	entityFields, _ := coreUtils.GetEntityFields(reflect.New(entityType).Interface())
	systemFields := coreUtils.GetSystemFields()

	fields := buildFieldInfo(entityType, systemFields, entityFields, sortField, sortDir)
	dataSlice := buildDataSlice(records, systemFields, entityFields)

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
