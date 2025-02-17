package querybuilder

import (
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
	tableName, fieldNames, err := qb.extractTableAndFields(tempEntity)
	if err != nil {
		println("Error1:", err)
		return err
	}

	query := fmt.Sprintf("SELECT %s FROM %s", joinFields(fieldNames), tableName)
	query = qb.prepareWhere(query)
	query = qb.prepareOrderBy(query)
	query = qb.prepareLimit(query)
	query = qb.prepareOffset(query)

	qb.query = query

	// Виконуємо запит
	rows, err := qb.client.GetDb().Query(qb.query, qb.args...)
	if err != nil {
		println("Error2:", err)
		return err
	}
	defer rows.Close()

	sliceValue.Elem().Set(reflect.MakeSlice(sliceValue.Elem().Type(), 0, 0))

	for rows.Next() {
		elem := reflect.New(elemType).Elem()
		fieldPointers := make([]interface{}, len(fieldNames))

		fmt.Println("fieldNames", fieldNames)
		for i := range fieldNames {
			if i == 0 {
				// TODO: implement good support for entities.Model fields
				fieldPointers[i] = &elem.Field(i).Addr().Interface().(*entities2.Model).ID
				continue
			}

			fieldPointers[i] = elem.Field(i).Addr().Interface()
		}

		if err := rows.Scan(fieldPointers...); err != nil {
			println("Error3:", err)
			return err
		}

		fmt.Println("fieldPointers:", fieldPointers)

		sliceValue.Elem().Set(reflect.Append(sliceValue.Elem(), elem))
	}

	return rows.Err()
}
