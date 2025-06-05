package adapter_mysql

import (
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
)

func (ap AdapterMySQL) GetFromSubquery(tableName string, where string, orderBy []string, limit int, offset int) string {
	return base.GetFromSubquery(tableName, where, orderBy, limit, offset)
}

func (ap AdapterMySQL) GetReadQuery(tableName string, fields []string, fromSubquery string) string {
	return base.GetReadQuery(tableName, fields, fromSubquery)
}

func (ap AdapterMySQL) GetSelectQuery(selectValue string, fromValue string) string {
	return base.GetSelectQuery(selectValue, fromValue)
}

func (ap AdapterMySQL) GetSelectWhereIn(tableName string, selectValue string, fieldName string, fieldValues []string) string {
	return base.GetSelectWhereIn(tableName, selectValue, fieldName, fieldValues)
}
