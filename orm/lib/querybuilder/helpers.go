package querybuilder

import (
	"errors"
	"github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/core/utils"
	"reflect"
)

// extractTableAndFields: Extracts table name and field names from an entity.
func (qb *QueryBuilder) extractTableAndFields(entity any, prefix bool) (string, []string, []string, error) {
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

	entityFieldNames := utils.ExtractFields(entityType)
	systemFieldNames := utils.ExtractFields(reflect.TypeOf(entities.Model{}))

	if !prefix {
		return tableName, entityFieldNames, systemFieldNames, nil
	}

	mappedEntityFields := utils.AddPrefix(tableName, entityFieldNames)
	mappedSystemFields := utils.AddPrefix(tableName, systemFieldNames)

	return tableName, mappedEntityFields, mappedSystemFields, nil
}

func (qb *QueryBuilder) extractTableAndFieldsFromType(elemType reflect.Type, prefix bool) (string, []string, []string, error) {
	tempEntity := reflect.New(elemType).Interface()
	return qb.extractTableAndFields(tempEntity, prefix)
}
