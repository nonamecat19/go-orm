package adapter_mssql

import (
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
)

func (ap AdapterMSSQL) DeleteFromTable(tableName string) string {
	return base.DeleteFromTable(tableName)
}
