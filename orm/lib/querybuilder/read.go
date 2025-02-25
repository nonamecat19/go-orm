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

	fields := append([]string{}, systemFieldNames...)
	fields = append(fields, entityFieldNames...)

	if err != nil {
		println("Error1:", err)
		return err
	}

	query := fmt.Sprintf("SELECT %s FROM %s", joinFields(fields), tableName)
	query = qb.prepareWhere(query)
	query = qb.prepareOrderBy(query)
	query = qb.prepareLimit(query)
	query = qb.prepareOffset(query)

	qb.query = query

	rows, err := qb.client.GetDb().Query(qb.query, qb.args...)
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	sliceValue.Elem().Set(reflect.MakeSlice(sliceValue.Elem().Type(), 0, 0))

	for rows.Next() {
		elem := reflect.New(elemType).Elem()
		var fieldPointers []interface{}

		fmt.Println(entityFieldNames)
		fmt.Println(systemFieldNames)

		model := elem.Field(0).Addr().Interface().(*entities2.Model)
		fieldPointers = append(fieldPointers, &model.ID)
		fieldPointers = append(fieldPointers, &model.CreatedAt)
		fieldPointers = append(fieldPointers, &model.UpdatedAt)
		fieldPointers = append(fieldPointers, &model.DeletedAt)

		for i := range entityFieldNames {
			fieldPointers = append(fieldPointers, elem.Field(i+1).Addr().Interface())
		}

		fmt.Println(fieldPointers)

		if err := rows.Scan(fieldPointers...); err != nil {
			fmt.Println("Error8:", err)
			return err
		}

		sliceValue.Elem().Set(reflect.Append(sliceValue.Elem(), elem))
	}

	return rows.Err()
}
