package querybuilder

import (
	"errors"
	"reflect"
	"strings"
)

// extractTableAndFields: Extracts table name and field names from an entity.
func (qb *QueryBuilder) extractTableAndFields(entity interface{}) (string, []string, error) {
	entityType := reflect.TypeOf(entity)
	if entityType.Kind() != reflect.Ptr || entityType.Elem().Kind() != reflect.Struct {
		return "", nil, errors.New("entity must be a pointer to a struct")
	}

	entityType = entityType.Elem()
	tableName := ""
	if tableNameMethod, ok := reflect.New(entityType).Interface().(interface{ TableName() string }); ok {
		tableName = tableNameMethod.TableName()
	} else {
		return "", nil, errors.New("entity struct must implement Info() string method")
	}

	var fieldNames []string
	for i := 0; i < entityType.NumField(); i++ {
		fieldTag := entityType.Field(i).Tag.Get("db")
		if fieldTag != "" {
			fieldNames = append(fieldNames, fieldTag)
		}
	}

	return tableName, fieldNames, nil
}

func joinFields(fields []string) string {
	return strings.Join(fields, ", ")
}
