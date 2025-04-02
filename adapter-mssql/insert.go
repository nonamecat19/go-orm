package adapter_mssql

import base "adapter-base"

func (ap AdapterMSSQL) Insert(tableName string, fieldNames []string, values []any, args []any) (string, []any) {
	return base.Insert(tableName, fieldNames, values, args)
}
