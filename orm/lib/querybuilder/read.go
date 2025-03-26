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

	err = qb.handleFindRows(sliceValue, elemType, rows)
	if err != nil {
		return err
	}

	err = qb.preloadRelations(sliceValue, elemType)
	if err != nil {
		return err
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
			fmt.Println("1")
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
		query = fmt.Sprintf("SELECT %s", JoinFieldsStrictly(fields))
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

func (qb *QueryBuilder) preloadRelationPointer(field reflect.StructField, sliceValue reflect.Value, elemType reflect.Type) error {
	structType := field.Type.Elem()

	infoInstance, ok := reflect.New(structType).Interface().(interface{ Info() string })

	if !ok {
		return fmt.Errorf("struct %s must implement Info() string method", structType.Name())
	}

	tableName := infoInstance.Info()
	relationTag := field.Tag.Get("relation")

	uniqueUserIDs := make(map[any]bool)
	var userIDs []any

	for i := 0; i < sliceValue.Elem().Len(); i++ {
		elem := sliceValue.Elem().Index(i)
		var foundField reflect.StructField

		t := elem.Type()
		for i := 0; i < t.NumField(); i++ {
			elemField := t.Field(i)
			if elemField.Tag.Get("db") == relationTag {
				foundField = elemField
				break
			}
		}

		if foundField.Name == "" {
			return fmt.Errorf("field with tag 'db:\"%s\"' not found for type %s", relationTag, elem.Type().Name())
		}

		elemField := elem.FieldByName(foundField.Name)
		if !elemField.IsValid() {
			return fmt.Errorf("field with tag 'db:\"%s\"' is not valid for type %s", relationTag, elem.Type().Name())
		}

		id := elemField.Interface()
		if _, exists := uniqueUserIDs[id]; !exists {
			uniqueUserIDs[id] = true
			userIDs = append(userIDs, id)
		}
	}

	stringUserIDs := make([]string, 0, len(userIDs))
	for _, id := range userIDs {
		val := reflect.ValueOf(id)

		if val.Kind() == reflect.Ptr {
			ptr, _ := id.(*int64)
			if ptr != nil {
				stringUserIDs = append(stringUserIDs, fmt.Sprintf("%v", val.Elem().Interface()))
			}
		} else {
			stringUserIDs = append(stringUserIDs, fmt.Sprintf("%v", id))
		}
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id IN (%s)", tableName, JoinFields(stringUserIDs))

	rows, err := qb.client.GetDb().Query(query)
	if err != nil {
		return fmt.Errorf("failed to execute query for %s preload: %w", relationTag, err)
	}
	defer rows.Close()

	relationStructFieldName, err := utils.GetFieldNameByTagValue(elemType, relationTag)

	if err != nil {
		return fmt.Errorf("failed to get field name by tag value: %w", err)
	}

	err = qb.handlePreloadPtr(sliceValue, field, rows, relationStructFieldName)
	if err != nil {
		return fmt.Errorf("failed to preload %s: %w", relationTag, err)
	}

	return nil
}

func (qb *QueryBuilder) handlePreloadPtr(sliceValue reflect.Value, field reflect.StructField, rows *sql.Rows, relationFieldName string) error {
	elemType := field.Type.Elem()

	tableName, entityFieldNames, systemFieldNames, err := qb.extractTableAndFieldsFromType(elemType)
	if err != nil {
		return err
	}

	rowMap := make(map[any]interface{})

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

		rowMap[model.ID] = elem.Interface()
	}

	for i := 0; i < sliceValue.Elem().Len(); i++ {
		currentElem := sliceValue.Elem().Index(i)

		relationId := currentElem.FieldByName(relationFieldName)

		var id int64
		idValue := reflect.ValueOf(relationId.Interface())
		if idValue.Kind() == reflect.Ptr {
			if !idValue.IsNil() {
				id = idValue.Elem().Interface().(int64)
			}
		} else {
			castedId, ok := relationId.Interface().(int64)
			if !ok {
				return fmt.Errorf("relationId is not of type int64 or *int64")
			}
			id = castedId
		}
		value := rowMap[id]

		relationField := currentElem.FieldByName(elemType.Name())

		if value != nil && relationField.IsValid() && relationField.CanSet() {
			if reflect.TypeOf(value).Kind() == reflect.Ptr {
				relationField.Set(reflect.ValueOf(value))
			} else {
				userPtr := reflect.New(reflect.TypeOf(value))
				userPtr.Elem().Set(reflect.ValueOf(value))
				relationField.Set(userPtr)
			}
		}
	}

	return nil
}
