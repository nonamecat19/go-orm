package querybuilder

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/utils"
	"reflect"
)

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
				newArg = value.Elem()
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
