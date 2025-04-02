package adapter_mssql

import (
	base "adapter-base"
)

func (ap AdapterMSSQL) Update(tableName string) string {
	return base.Update(tableName)
}
