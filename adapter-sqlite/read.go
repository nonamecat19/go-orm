package adapter_sqlite

import (
	base "adapter-base"
)

func (ap AdapterSQLite) GetFromSubquery(tableName string, where string, orderBy []string, limit int, offset int) string {
	return base.GetFromSubquery(tableName, where, orderBy, limit, offset)
}

func (ap AdapterSQLite) GetReadQuery(tableName string, fields []string, fromSubquery string) string {
	return base.GetReadQuery(tableName, fields, fromSubquery)
}

func (ap AdapterSQLite) GetSelectQuery(selectValue string, fromValue string) string {
	return base.GetSelectQuery(selectValue, fromValue)
}

func (ap AdapterSQLite) GetSelectWhereIn(tableName string, selectValue string, fieldName string, fieldValues []string) string {
	return base.GetSelectWhereIn(tableName, selectValue, fieldName, fieldValues)
}
