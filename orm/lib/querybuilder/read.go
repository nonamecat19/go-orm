package querybuilder

import (
	"database/sql"
	"fmt"
	entities2 "github.com/nonamecat19/go-orm/core/lib/entities"
	"reflect"
)

// FindOne initialized a SELECT query for one record of the specified entity
func (qb *QueryBuilder) FindOne() {
}

// FindMany initializes a SELECT query for the specified entity.
func (qb *QueryBuilder) FindMany(entities interface{}) error {
	sliceValue := reflect.ValueOf(entities)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("entities must be a pointer to a slice")
	}

	elemType := sliceValue.Elem().Type().Elem()

	tempEntity := reflect.New(elemType).Interface()
	tableName, entityFieldNames, systemFieldNames, err := qb.extractTableAndFields(tempEntity)
	if err != nil {
		return err
	}

	fields := append(systemFieldNames, entityFieldNames...)

	query := fmt.Sprintf("SELECT %s FROM %s", joinFields(fields), tableName)
	query = qb.prepareWhere(query)
	query = qb.prepareOrderBy(query)
	query = qb.prepareLimit(query)
	query = qb.prepareOffset(query)

	qb.query = query

	rows, err := qb.Query()

	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	sliceValue.Elem().Set(reflect.MakeSlice(sliceValue.Elem().Type(), 0, 0))

	elem := reflect.New(elemType).Elem()

	model := elem.Field(0).Addr().Interface().(*entities2.Model)
	systemFieldsMap := map[string]any{
		"id":         &model.ID,
		"created_at": &model.CreatedAt,
		"updated_at": &model.UpdatedAt,
		"deleted_at": &model.DeletedAt,
	}

	for rows.Next() {
		var fieldPointers []interface{}

		for _, name := range systemFieldNames {
			ptr := systemFieldsMap[name]
			if ptr != nil {
				fieldPointers = append(fieldPointers, ptr)
			}
		}

		for i := range entityFieldNames {
			fieldPointers = append(fieldPointers, elem.Field(i+1).Addr().Interface())
		}

		if err := rows.Scan(fieldPointers...); err != nil {
			return err
		}

		sliceValue.Elem().Set(reflect.Append(sliceValue.Elem(), elem))
	}

	return rows.Err()
}
