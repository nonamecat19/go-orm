package querybuilder

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/utils"
	"reflect"
)

func (qb *QueryBuilder) InsertOne(entity interface{}) error {

	return nil
}

func (qb *QueryBuilder) InsertMany(entities interface{}) error {
	sliceValue := reflect.ValueOf(entities)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("entities must be a pointer to a slice")
	}

	elementType := sliceValue.Elem().Type().Elem()

	tableName, entityFieldNames, systemFieldNames, err := qb.extractTableAndFieldsFromType(elementType, false)

	slice := sliceValue.Elem()

	var stringRecords []string
	var queryArgs []interface{}

	for i := 0; i < slice.Len(); i++ {

		entity := slice.Index(i)

		for _, entityFieldName := range entityFieldNames {
			currentFieldName, err := utils.GetFieldNameByTagValue(entity.Type(), entityFieldName)

			if err != nil {
				return err
			}

			value := entity.FieldByName(currentFieldName)

			switch value.Kind() {
			case reflect.Ptr:
				if value.IsNil() {
					queryArgs = append(queryArgs, "NULL")
				} else {
					queryArgs = append(queryArgs, value.Elem())
				}
			default:
				queryArgs = append(queryArgs, value.Interface())
			}

			fmt.Println(entity)
		}

		stringRecords = append(stringRecords, fmt.Sprintf("(%s)", JoinFields(utils.GenerateParamsSlice(len(entityFieldNames)))))
	}

	fmt.Println(tableName, entityFieldNames, systemFieldNames, err)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;", tableName, JoinFields(entityFieldNames), JoinFields(stringRecords))
	qb.query = qb.normalizeSqlWithArgs(query)
	qb.args = append(qb.args, queryArgs...)

	rows, err := qb.ExecuteBuilderQuery()
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
