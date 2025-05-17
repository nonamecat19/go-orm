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
func (qb *QueryBuilder) FindMany(entities any) error {
	sliceValue := reflect.ValueOf(entities)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("entities must be a pointer to a slice")
	}

	err := qb.prepareFindQuery(sliceValue.Elem().Type().Elem())
	if err != nil {
		return err
	}

	rows, err := qb.ExecuteBuilderQuery()
	if err != nil {
		return err
	}
	defer rows.Close()

	err = qb.handleFindRows(sliceValue, sliceValue.Elem().Type().Elem(), rows)
	if err != nil {
		return err
	}

	err = qb.preloadRelations(sliceValue, sliceValue.Elem().Type().Elem())
	if err != nil {
		return err
	}

	return rows.Err()
}

func (qb *QueryBuilder) handleFindRows(sliceValue reflect.Value, elemType reflect.Type, rows *sql.Rows) error {
	tableName, entityFieldNames, systemFieldNames, err := utils.ExtractTableAndFieldsFromType(elemType, true)
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
		var fieldPointers []any

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
				fieldPointers = append(fieldPointers, new(any))
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
	tableName, entityFieldNames, systemFieldNames, err := utils.ExtractTableAndFieldsFromType(elemType, true)
	if err != nil {
		return err
	}

	fields := append(systemFieldNames, entityFieldNames...)

	if len(qb.selectFields) > 0 {
		fields = utils.StringsIntersection(fields, qb.selectFields)
	}

	for _, join := range qb.joins {
		fields = append(fields, join.Select...)
	}

	fromSubquery := qb.adapter.GetFromSubquery(tableName, qb.where, qb.orderBy, qb.limit, qb.offset)
	query := qb.adapter.GetReadQuery(tableName, fields, fromSubquery)

	qb.query = qb.adapter.PrepareJoins(query, qb.joins)

	return nil
}

func (qb *QueryBuilder) preloadRelations(sliceValue reflect.Value, elemType reflect.Type) error {
	var err error
	if len(qb.preloads) == 0 {
		return nil
	}

	elem := reflect.New(elemType).Elem()
	t := reflect.TypeOf(elem.Interface())

	for _, preload := range qb.preloads {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			dbTag := field.Tag.Get("db")
			if dbTag == "" || dbTag != preload {
				continue
			}

			fieldKind := field.Type.Kind()

			if fieldKind == reflect.Ptr {
				err = qb.preloadRelationPointer(field, sliceValue, elemType)
			} else if fieldKind == reflect.Slice {
				err = qb.preloadRelationSlice(field, sliceValue, elemType)
			} else {
				return fmt.Errorf("%s must be relation field", preload)
			}

			if err != nil {
				return err
			}
		}
	}

	return nil
}
