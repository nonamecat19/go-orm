package querybuilder

import (
	"errors"
	"fmt"
	"github.com/nonamecat19/go-orm/core/lib/entities"
	"reflect"
	"strings"
)

// addPrefix adds a prefix to each string in the slice
func addPrefix(prefix string, slice []string) []string {
	result := make([]string, len(slice))
	for i, s := range slice {
		result[i] = fmt.Sprintf("%s.%s", prefix, s)
	}
	return result
}

// extractFields extract all field names from entity
func extractFields(entity reflect.Type) []string {
	var fieldNames []string

	for i := 0; i < entity.NumField(); i++ {
		fieldTags := entity.Field(i).Tag
		dbTag := fieldTags.Get("db")
		relationTag := fieldTags.Get("relation")

		if dbTag != "" && relationTag == "" {
			fieldNames = append(fieldNames, dbTag)
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

	mappedEntityFields := addPrefix(tableName, entityFieldNames)
	mappedSystemFields := addPrefix(tableName, systemFieldNames)

	return tableName, mappedEntityFields, mappedSystemFields, nil
}

func (qb *QueryBuilder) extractTableAndFieldsFromType(elemType reflect.Type) (string, []string, []string, error) {
	tempEntity := reflect.New(elemType).Interface()
	return qb.extractTableAndFields(tempEntity)
}

func JoinFields(fields []string) string {
	return strings.Join(fields, ", ")
}

func JoinFieldsStrictly(fields []string) string {
	mappedFields := make([]string, len(fields))
	for i, field := range fields {
		mappedFields[i] = fmt.Sprintf("%s AS \"%s\"", field, field)
	}
	return strings.Join(mappedFields, ", ")
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func getModelFields(model interface{}) map[string]any {
	v := reflect.ValueOf(model).Elem()
	t := v.Type()

	fields := make(map[string]any)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dbTag := field.Tag.Get("db")

		fieldPtr := v.Field(i).Addr().Interface()

		if dbTag != "" {
			fields[dbTag] = fieldPtr
		}
	}

	return fields
}

// normalizeSqlWithArgs change "?" to database valid syntax
func (qb *QueryBuilder) normalizeSqlWithArgs(sql string) string {
	placeholderIndex := len(qb.args) + 1

	for {
		placeholder := fmt.Sprintf("$%d", placeholderIndex)
		sql = strings.Replace(sql, "?", placeholder, 1) // Replace only the first '?' occurrence
		if !strings.Contains(sql, "?") {
			break
		}
		placeholderIndex++
	}

	return sql
}
