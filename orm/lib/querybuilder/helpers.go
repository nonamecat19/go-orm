package querybuilder

import (
	"errors"
	"github.com/nonamecat19/go-orm/core/lib/entities"
	"reflect"
	"strings"
)

func extractFields(entity reflect.Type) []string {
	var fieldNames []string

	for i := 0; i < entity.NumField(); i++ {
		fieldTag := entity.Field(i).Tag.Get("db")
		if fieldTag != "" {
			fieldNames = append(fieldNames, fieldTag)
		}
	}

	return fieldNames
}

// extractTableAndFields: Extracts table name and field names from an entity.
func (qb *QueryBuilder) extractTableAndFields(entity interface{}) (string, []string, []string, error) {
	entityType := reflect.TypeOf(entity)
	if entityType.Kind() != reflect.Ptr || entityType.Elem().Kind() != reflect.Struct {
		return "", nil, nil, errors.New("entity must be a pointer to a struct")
	}

	entityType = entityType.Elem()
	tableName := ""
	if tableNameMethod, ok := reflect.New(entityType).Interface().(interface{ Info() string }); ok {
		tableName = tableNameMethod.Info()
	} else {
		return "", nil, nil, errors.New("entity struct must implement Info() string method")
	}

	entityFieldNames := extractFields(entityType)
	systemFieldNames := extractFields(reflect.TypeOf(entities.Model{}))

	return tableName, entityFieldNames, systemFieldNames, nil
}

func joinFields(fields []string) string {
	return strings.Join(fields, ", ")
}
