package adapter_base

import "fmt"

func GetReadQuery(tableName string, fields []string, fromSubquery string) string {
	return fmt.Sprintf("SELECT %s FROM (%s) AS %s", JoinFieldsStrictly(fields), fromSubquery, tableName)
}

func GetFromSubquery(tableName string, where string, orderBy []string, limit int, offset int) string {
	fromQuery := fmt.Sprintf("SELECT * FROM %s", tableName)
	fromQuery = PrepareWhere(fromQuery, where)
	fromQuery = PrepareOrderBy(fromQuery, orderBy)
	fromQuery = PrepareLimit(fromQuery, limit)
	fromQuery = PrepareOffset(fromQuery, offset)

	return fromQuery
}
