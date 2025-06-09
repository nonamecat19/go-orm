package querybuilder

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/utils"
	"reflect"
)

func (qb *QueryBuilder) InsertMap(entity any, mapFields map[string]any) error {
	elementType := reflect.TypeOf(entity)
	if elementType.Kind() != reflect.Struct {
		return fmt.Errorf("entity must be a struct")
	}

	tableName, entityFieldNames, _, err := utils.ExtractTableAndFieldsFromType(elementType, false)
	if err != nil {
		return err
	}

	var fieldNames []string
	var queryArgs []any

	for key, value := range mapFields {
		if !utils.Contains(entityFieldNames, key) {
			return fmt.Errorf("field %s does not exist in entity", key)
		}
		fieldNames = append(fieldNames, key)
		queryArgs = append(queryArgs, value)
	}

	if len(fieldNames) == 0 {
		return fmt.Errorf("no valid fields provided in mapFields")
	}

	qb.query, qb.args = qb.adapter.Insert(tableName, fieldNames, queryArgs, qb.args)

	rows, err := qb.ExecuteBuilderQuery()
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (qb *QueryBuilder) InsertOne(entity any) error {
	elementValue := reflect.ValueOf(entity)
	if elementValue.Kind() != reflect.Struct {
		return fmt.Errorf("entities must be a struct")
	}

	elementType := elementValue.Type()
	sliceType := reflect.SliceOf(elementType)
	sliceValue := reflect.MakeSlice(sliceType, 0, 1)
	sliceValue = reflect.Append(sliceValue, elementValue)

	return qb.insertSlice(sliceValue)
}

func (qb *QueryBuilder) InsertMany(entities any) error {
	sliceValue := reflect.ValueOf(entities)
	if sliceValue.Kind() != reflect.Slice {
		return fmt.Errorf("entities must be a slice")
	}

	return qb.insertSlice(sliceValue)
}

func (qb *QueryBuilder) insertSlice(sliceValue reflect.Value) error {
	elementType := sliceValue.Type().Elem()

	tableName, entityFieldNames, _, err := utils.ExtractTableAndFieldsFromType(elementType, false)

	if err != nil {
		return err
	}

	var queryArgs []any

	for i := 0; i < sliceValue.Len(); i++ {
		entity := sliceValue.Index(i)

		for _, entityFieldName := range entityFieldNames {
			currentFieldName, err := utils.GetFieldNameByTagValue(entity.Type(), entityFieldName)
			if err != nil {
				return err
			}

			var newArg any
			value := entity.FieldByName(currentFieldName)

			switch value.Kind() {
			case reflect.Ptr:
				if value.IsNil() {
					newArg = nil
				} else {
					newArg = value.Elem().Interface()
				}
			default:
				newArg = value.Interface()
			}

			queryArgs = append(queryArgs, newArg)
		}
	}

	qb.query, qb.args = qb.adapter.Insert(tableName, entityFieldNames, queryArgs, qb.args)

	rows, err := qb.ExecuteBuilderQuery()
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
