package adapter_postgres

import (
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
)

func (ap AdapterPostgres) GetFromSubquery(tableName string, where string, orderBy []string, limit int, offset int) string {
	return base.GetFromSubquery(tableName, where, orderBy, limit, offset)
}

func (ap AdapterPostgres) GetReadQuery(tableName string, fields []string, fromSubquery string) string {
	return base.GetReadQuery(tableName, fields, fromSubquery)
}

func (ap AdapterPostgres) GetSelectQuery(selectValue string, fromValue string) string {
	return base.GetSelectQuery(selectValue, fromValue)
}

func (ap AdapterPostgres) GetSelectWhereIn(tableName string, selectValue string, fieldName string, fieldValues []string) string {
	return base.GetSelectWhereIn(tableName, selectValue, fieldName, fieldValues)
}
