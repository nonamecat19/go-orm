package querybuilder

import (
	"errors"
	"fmt"
	"reflect"
)

// FindOne initialized a SELECT query for one record of the specified entity
func (qb *QueryBuilder) FindOne() {
}

// FindMany initializes a SELECT query for the specified entity.
func (qb *QueryBuilder) FindMany(entity interface{}) (*string, error) {
	entityType := reflect.TypeOf(entity)
	if entityType.Kind() != reflect.Ptr || entityType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("entity must be a pointer to a struct")
	}

	entityType = entityType.Elem()
	tableName := ""
	if tableNameMethod, ok := reflect.New(entityType).Interface().(interface{ TableName() string }); ok {
		tableName = tableNameMethod.TableName()
	} else {
		return nil, errors.New("entity struct must implement TableName() string method")
	}

	var fieldNames []string
	for i := 0; i < entityType.NumField(); i++ {
		fieldTag := entityType.Field(i).Tag.Get("db")
		if fieldTag != "" {
			fieldNames = append(fieldNames, fieldTag)
		}
	}

	query := fmt.Sprintf("SELECT %s FROM %s", joinFields(fieldNames), tableName)

	query = qb.prepareWhere(query)
	query = qb.prepareOrderBy(query)
	query = qb.prepareLimit(query)
	query = qb.prepareOffset(query)

	qb.query = query
	return &query, nil
}
