package adapter_sqlite

import base "adapter-base"

func (ap AdapterSQLite) Insert(tableName string, fieldNames []string, values []any, args []any) (string, []any) {
	return base.Insert(tableName, fieldNames, values, args)
}
