package querybuilder

import (
	"errors"
	"fmt"
	"github.com/nonamecat19/go-orm/core/utils"
	"reflect"
)

// UpdateMany initializes an UPDATE query for the specified entity.
func (qb *QueryBuilder) UpdateMany(entity any) error {
	tableName, _, _, err := utils.ExtractTableAndFields(entity, true)
	if err != nil {
		return err
	}

	if len(qb.where) == 0 {
		return errors.New("UPDATE query requires at least one WHERE clause to prevent accidental update of all rows")
	}

	if len(qb.set) == 0 {
		return errors.New("no updates provided")
	}

	query := qb.adapter.Update(tableName)
	query = qb.prepareSet(query)
	query = qb.prepareWhere(query)

	qb.query = query

	_, err = qb.ExecuteBuilderQuery()

	return err
}

func (qb *QueryBuilder) SetValues(values map[string]any) *QueryBuilder {
	qb.set = values
	return qb
}

func (qb *QueryBuilder) UpdateMap(entity any, mapFields map[string]any) error {
	elementType := reflect.TypeOf(entity)
	if elementType.Kind() != reflect.Struct {
		return fmt.Errorf("entity must be a struct")
	}

	tableName, entityFieldNames, _, err := utils.ExtractTableAndFieldsFromType(elementType, false)
	if err != nil {
		return err
	}

	for key := range mapFields {
		if !utils.Contains(entityFieldNames, key) {
			return fmt.Errorf("field %s does not exist in entity", key)
		}
	}

	if len(mapFields) == 0 {
		return fmt.Errorf("no valid fields provided in mapFields")
	}

	if len(qb.where) == 0 {
		return errors.New("UPDATE query requires at least one WHERE clause to prevent accidental update of all rows")
	}

	qb.set = mapFields

	query := qb.adapter.Update(tableName)
	query = qb.prepareSet(query)
	query = qb.prepareWhere(query)

	qb.query = query

	_, err = qb.ExecuteBuilderQuery()

	return err
}
