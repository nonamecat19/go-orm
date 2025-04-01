package querybuilder

import (
	"database/sql"
	"fmt"
	entities2 "github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/core/utils"
	"reflect"
)

func (qb *QueryBuilder) preloadRelationSlice(field reflect.StructField, sliceValue reflect.Value, elemType reflect.Type) error {
	structType := field.Type.Elem()

	infoInstance, ok := reflect.New(structType).Interface().(interface{ Info() string })

	if !ok {
		return fmt.Errorf("struct %s must implement Info() string method", structType.Name())
	}

	tableName := infoInstance.Info()
	relationTag := field.Tag.Get("relation")

	uniqueEntityIDs := make(map[any]bool)
	var entityIDs []any

	for i := 0; i < sliceValue.Elem().Len(); i++ {
		elem := sliceValue.Elem().Index(i)
		// to get Entity -> Model -> ID
		id := elem.Field(0).Field(0).Interface()
		if _, exists := uniqueEntityIDs[id]; exists {
			continue
		}
		uniqueEntityIDs[id] = true
		entityIDs = append(entityIDs, id)
	}

	stringEntityIDs := make([]string, 0, len(entityIDs))
	for _, id := range entityIDs {
		val := reflect.ValueOf(id)

		if val.Kind() == reflect.Ptr {
			ptr, ok := id.(*int64)
			if !ok {
				return fmt.Errorf("id is not of type int64 or *int64")
			}
			if ptr != nil {
				stringEntityIDs = append(stringEntityIDs, fmt.Sprintf("%v", val.Elem().Interface()))
			}
		} else {
			stringEntityIDs = append(stringEntityIDs, fmt.Sprintf("%v", id))
		}
	}

	_, entityFieldNames, systemFieldNames, err := qb.extractTableAndFieldsFromType(field.Type.Elem(), true)
	if err != nil {
		return err
	}

	fields := append(systemFieldNames, entityFieldNames...)

	//if len(qb.selectFields) > 0 {
	//	fields = utils.StringsIntersection(fields, qb.selectFields)
	//}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (%s)", qb.adapter.JoinFieldsStrictly(fields), tableName, relationTag, qb.adapter.JoinFields(stringEntityIDs))

	rows, err := qb.ExecuteQuery(query)
	if err != nil {
		return fmt.Errorf("failed to execute query for %s preload: %w", relationTag, err)
	}
	defer rows.Close()

	relationStructFieldName, err := utils.GetFieldNameByTagValue(elemType, tableName)
	if err != nil {
		return fmt.Errorf("failed to get field name by tag value: %w", err)
	}

	relationTagField, err := utils.GetFieldNameByTagValue(structType, relationTag)
	if err != nil {
		return fmt.Errorf("failed to get field name: %s", relationTag)
	}

	err = qb.handlePreloadSlice(sliceValue, field, rows, relationStructFieldName, relationTagField)
	if err != nil {
		return fmt.Errorf("failed to preload %s: %w", relationTag, err)
	}

	return nil
}

func (qb *QueryBuilder) handlePreloadSlice(sliceValue reflect.Value, field reflect.StructField, rows *sql.Rows, relationFieldName string, relationFieldTag string) error {
	elemType := field.Type.Elem()

	tableName, entityFieldNames, systemFieldNames, err := qb.extractTableAndFieldsFromType(elemType, true)
	if err != nil {
		return err
	}

	rowMap := make(map[any]any)

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

		systemFields := utils.StringsIntersection(systemFieldNames, fields)
		for _, name := range systemFields {
			ptr := systemFieldsMap[name]
			if ptr != nil {
				fieldPointers = append(fieldPointers, ptr)
			}
		}

		fieldMap := make(map[string]int)
		t := reflect.TypeOf(elem.Interface())

		for i := 0; i < t.NumField(); i++ {
			dbTag, ok := t.Field(i).Tag.Lookup("db")
			if ok {
				fieldMap[fmt.Sprintf("%s.%s", tableName, dbTag)] = i
			}
		}

		bodyFields := utils.StringsIntersection(entityFieldNames, fields)
		for _, fieldName := range bodyFields {
			idx, exists := fieldMap[fieldName]
			if !exists {
				continue
			}

			currentField := elem.Field(idx)
			if currentField.Kind() == reflect.Ptr {
				if currentField.IsNil() {
					newValue := reflect.New(currentField.Type().Elem())
					currentField.Set(newValue)
				}
				fieldPointers = append(fieldPointers, currentField.Interface())

			} else {
				fieldPointers = append(fieldPointers, currentField.Addr().Interface())
			}
		}

		if err := rows.Scan(fieldPointers...); err != nil {
			return err
		}

		relationFieldValue := reflect.ValueOf(elem.Interface()).FieldByName(relationFieldTag)
		var key int64
		var ok bool

		if relationFieldValue.Kind() == reflect.Ptr {
			key, ok = relationFieldValue.Elem().Interface().(int64)
		} else {
			key, ok = relationFieldValue.Interface().(int64)
		}

		if !ok {
			return fmt.Errorf("relationFieldValue is not of type int64 or *int64")
		}

		if existingSlice, ok := rowMap[key]; ok {
			rowMap[key] = append(existingSlice.([]any), elem.Interface())
		} else {
			rowMap[key] = []any{elem.Interface()}
		}
	}

	for i := 0; i < sliceValue.Elem().Len(); i++ {
		currentElem := sliceValue.Elem().Index(i)

		entitySystemFields, ok := currentElem.Field(0).Interface().(entities2.Model)
		if !ok {
			return fmt.Errorf("failed to cast to entities.Model")
		}

		id := entitySystemFields.ID
		value := rowMap[id]

		relationField := currentElem.FieldByName(relationFieldName)

		if value == nil || !relationField.IsValid() || !relationField.CanSet() {
			continue
		}
		entityPtr := reflect.New(reflect.TypeOf(value))
		entityPtr.Elem().Set(reflect.ValueOf(value))

		relationFieldType := relationField.Type().Elem()

		var convertedSlice reflect.Value

		if relationField.Kind() != reflect.Slice {
			return fmt.Errorf("relationField is not a slice")
		}

		convertedSlice = reflect.MakeSlice(relationField.Type(), 0, len(value.([]any)))

		for _, v := range value.([]any) {
			entityValue := reflect.ValueOf(v)

			if !entityValue.Type().AssignableTo(relationFieldType) {
				return fmt.Errorf("incompatible type: %v cannot be assigned to %v", entityValue.Type(), relationFieldType)
			}
			convertedSlice = reflect.Append(convertedSlice, entityValue)
		}
		relationField.Set(convertedSlice)
	}

	return nil
}
