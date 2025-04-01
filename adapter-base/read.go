package adapter_base

import "fmt"

func GetReadQuery(tableName string, fields []string, fromSubquery string) string {
	selectQuery := GetSelectQuery(JoinFieldsStrictly(fields), fmt.Sprintf("(%s)", fromSubquery))
	return fmt.Sprintf("%s AS %s", selectQuery, tableName)
}

func GetFromSubquery(tableName string, where string, orderBy []string, limit int, offset int) string {
	fromQuery := GetSelectQuery("*", tableName)
	fromQuery = PrepareWhere(fromQuery, where)
	fromQuery = PrepareOrderBy(fromQuery, orderBy)
	fromQuery = PrepareLimit(fromQuery, limit)
	fromQuery = PrepareOffset(fromQuery, offset)

	return fromQuery
}

func GetSelectQuery(selectValue string, fromValue string) string {
	return fmt.Sprintf("SELECT %s FROM %s", selectValue, fromValue)
}

func GetSelectWhereIn(tableName string, fields []string, fieldName string, fieldValues []string) string {
	selectQuery := GetSelectQuery(JoinFieldsStrictly(fields), tableName)
	return fmt.Sprintf("%s WHERE %s IN (%s)", selectQuery, fieldName, JoinFields(fieldValues))
}
