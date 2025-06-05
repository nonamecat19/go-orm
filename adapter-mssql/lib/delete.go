package adapter_mssql

import (
	base "adapter-base/lib"
)

func (ap AdapterMSSQL) DeleteFromTable(tableName string) string {
	return base.DeleteFromTable(tableName)
}
