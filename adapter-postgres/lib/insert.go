package adapter_postgres

import (
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
)

func (ap AdapterPostgres) Insert(tableName string, fieldNames []string, values []any, args []any) (string, []any) {
	return base.Insert(tableName, fieldNames, values, args)
}
