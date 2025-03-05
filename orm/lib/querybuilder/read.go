package querybuilder

import (
	"database/sql"
	"fmt"
	entities2 "github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/core/utils"
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
	fields = utils.StringsIntersection(fields, qb.selectFields)

	var query string

	if len(fields) > 0 {
		query = fmt.Sprintf("SELECT %s", joinFields(fields))
	} else {
		query = "SELECT *"
	}

	for _, join := range qb.joins {
		for _, field := range join.Select {
			query += fmt.Sprintf(", %s", field)
		}
	}

	query += fmt.Sprintf(" FROM %s", tableName)

	for _, join := range qb.joins {
		query += fmt.Sprintf(" %s JOIN %s ON %s", join.JoinType, join.Table, join.Condition)
	}

	query = qb.prepareWhere(query)
	query = qb.prepareOrderBy(query)
	query = qb.prepareLimit(query)
	query = qb.prepareOffset(query)

	qb.query = query

	rows, err := qb.ExecuteQuery()
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
		tableName + ".id":         &model.ID,
		tableName + ".created_at": &model.CreatedAt,
		tableName + ".updated_at": &model.UpdatedAt,
		tableName + ".deleted_at": &model.DeletedAt,
	}

	for rows.Next() {
		var fieldPointers []interface{}

		for _, name := range utils.StringsIntersection(systemFieldNames, qb.selectFields) {
			ptr := systemFieldsMap[name]
			if ptr != nil {
				fieldPointers = append(fieldPointers, ptr)
			}
		}

		for i := range utils.StringsIntersection(entityFieldNames, qb.selectFields) {
			fieldPointers = append(fieldPointers, elem.Field(i+1).Addr().Interface())
		}

		for _, join := range qb.joins {
			for range join.Select {
				fieldPointers = append(fieldPointers, new(interface{}))
			}
		}

		if err := rows.Scan(fieldPointers...); err != nil {
			return err
		}

		sliceValue.Elem().Set(reflect.Append(sliceValue.Elem(), elem))
	}

	return rows.Err()
}
