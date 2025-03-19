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

	err := qb.prepareFindQuery(elemType)
	if err != nil {
		return err
	}

	rows, err := qb.ExecuteQuery()
	if err != nil {
		return err
	}
	defer rows.Close()

	err2 := qb.handleFindRows(sliceValue, elemType, rows)
	if err2 != nil {
		return err2
	}

	return rows.Err()
}

func (qb *QueryBuilder) handleFindRows(sliceValue reflect.Value, elemType reflect.Type, rows *sql.Rows) error {
	tableName, entityFieldNames, systemFieldNames, err := qb.extractTableAndFieldsFromType(elemType)
	if err != nil {
		return err
	}

	sliceValue.Elem().Set(reflect.MakeSlice(sliceValue.Elem().Type(), 0, 0))

	elem := reflect.New(elemType).Elem()

	model := elem.Field(0).Addr().Interface().(*entities2.Model)
	systemFieldsMap := map[string]any{
		tableName + ".id":         &model.ID,
		tableName + ".created_at": &model.CreatedAt,
		tableName + ".updated_at": &model.UpdatedAt,
		tableName + ".deleted_at": &model.DeletedAt,
	}

	fields := append(systemFieldNames, entityFieldNames...)

	if len(qb.selectFields) > 0 {
		fields = utils.StringsIntersection(fields, qb.selectFields)
	}

	for rows.Next() {
		var fieldPointers []interface{}

		for _, name := range utils.StringsIntersection(systemFieldNames, fields) {
			ptr := systemFieldsMap[name]
			if ptr != nil {
				fieldPointers = append(fieldPointers, ptr)
			}
		}

		fieldMap := make(map[string]int)
		t := reflect.TypeOf(elem.Interface())

		for i := 0; i < t.NumField(); i++ {
			if dbTag, ok := t.Field(i).Tag.Lookup("db"); ok {
				fieldMap[fmt.Sprintf("%s.%s", tableName, dbTag)] = i
			}
		}

		for _, fieldName := range utils.StringsIntersection(entityFieldNames, fields) {
			if idx, exists := fieldMap[fieldName]; exists {
				fieldPointers = append(fieldPointers, elem.Field(idx).Addr().Interface())
			}
		}

		for _, join := range qb.joins {
			for range join.Select {
				fieldPointers = append(fieldPointers, new(interface{}))
			}
		}

		if err := rows.Scan(fieldPointers...); err != nil {
			return err
		}

		results := reflect.Append(sliceValue.Elem(), elem)
		sliceValue.Elem().Set(results)
	}

	return nil
}

func (qb *QueryBuilder) prepareFindQuery(elemType reflect.Type) error {
	tableName, entityFieldNames, systemFieldNames, err := qb.extractTableAndFieldsFromType(elemType)
	if err != nil {
		return err
	}

	fields := append(systemFieldNames, entityFieldNames...)

	if len(qb.selectFields) > 0 {
		fields = utils.StringsIntersection(fields, qb.selectFields)
	}

	var query string

	for _, join := range qb.joins {
		for _, field := range join.Select {
			fields = append(fields, field)
		}
	}

	if len(fields) > 0 {
		query = fmt.Sprintf("SELECT %s", joinFieldsStrictly(fields))
	} else {
		query = "SELECT *"
	}

	fromQuery := fmt.Sprintf("SELECT * FROM %s", tableName)
	fromQuery = qb.prepareWhere(fromQuery)
	fromQuery = qb.prepareOrderBy(fromQuery)
	fromQuery = qb.prepareLimit(fromQuery)
	fromQuery = qb.prepareOffset(fromQuery)

	query += fmt.Sprintf(" FROM (%s) AS %s", fromQuery, tableName)

	for _, join := range qb.joins {
		query += fmt.Sprintf(" %s JOIN %s ON %s", join.JoinType, join.Table, join.Condition)
	}

	qb.query = query

	return nil
}
