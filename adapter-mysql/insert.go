package adapter_mysql

import base "adapter-base"

func (ap AdapterMySQL) Insert(tableName string, fieldNames []string, values []any, args []any) (string, []any) {
	return base.Insert(tableName, fieldNames, values, args)
}
