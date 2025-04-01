package adapter_postgres

import base "adapter-base"

func (ap AdapterPostgres) Insert(tableName string, fieldNames []string, values []any, args []any) (string, []any) {
	return base.Insert(tableName, fieldNames, values, args)
}
