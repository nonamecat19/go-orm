package adapter_postgres

import (
	base "adapter-base"
)

func (ap AdapterPostgres) GetFromSubquery(tableName string, where string, orderBy []string, limit int, offset int) string {
	return base.GetFromSubquery(tableName, where, orderBy, limit, offset)
}

func (ap AdapterPostgres) GetReadQuery(tableName string, fields []string, fromSubquery string) string {
	return base.GetReadQuery(tableName, fields, fromSubquery)
}
