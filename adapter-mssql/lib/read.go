package adapter_mssql

import (
	base "adapter-base/lib"
)

func (ap AdapterMSSQL) GetFromSubquery(tableName string, where string, orderBy []string, limit int, offset int) string {
	fromQuery := ap.GetSelectQuery("*", tableName)
	fromQuery = ap.PrepareWhere(fromQuery, where)
	fromQuery = ap.PrepareOrderBy(fromQuery, orderBy)
	fromQuery = ap.PrepareOffset(fromQuery, offset)
	fromQuery = ap.PrepareLimit(fromQuery, limit)

	return fromQuery
}

func (ap AdapterMSSQL) GetReadQuery(tableName string, fields []string, fromSubquery string) string {
	return base.GetReadQuery(tableName, fields, fromSubquery)
}

func (ap AdapterMSSQL) GetSelectQuery(selectValue string, fromValue string) string {
	return base.GetSelectQuery(selectValue, fromValue)
}

func (ap AdapterMSSQL) GetSelectWhereIn(tableName string, selectValue string, fieldName string, fieldValues []string) string {
	return base.GetSelectWhereIn(tableName, selectValue, fieldName, fieldValues)
}
