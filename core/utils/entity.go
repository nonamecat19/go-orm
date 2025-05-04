package utils

import (
	"errors"
	"github.com/nonamecat19/go-orm/core/lib/entities"
	"reflect"
)

func GetSystemFields() []string {
	return ExtractFields(reflect.TypeOf(entities.Model{}))
}

func GetEntityFields(entity any) ([]string, error) {
	entityType := reflect.TypeOf(entity)
	if entityType.Kind() != reflect.Ptr || entityType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("entity must be a pointer to a struct")
	}

	return ExtractFields(entityType.Elem()), nil
}

// ExtractTableAndFields Extracts table name and field names from an entity.
func ExtractTableAndFields(entity any, prefix bool) (string, []string, []string, error) {
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

	entityFieldNames := ExtractFields(entityType)
	systemFieldNames := GetSystemFields()

	if !prefix {
		return tableName, entityFieldNames, systemFieldNames, nil
	}

	mappedEntityFields := AddPrefix(tableName, entityFieldNames)
	mappedSystemFields := AddPrefix(tableName, systemFieldNames)

	return tableName, mappedEntityFields, mappedSystemFields, nil
}

func ExtractTableAndFieldsFromType(elemType reflect.Type, prefix bool) (string, []string, []string, error) {
	tempEntity := reflect.New(elemType).Interface()
	return ExtractTableAndFields(tempEntity, prefix)
}
