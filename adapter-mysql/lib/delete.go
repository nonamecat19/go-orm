package adapter_mysql

import (
	base "adapter-base/lib"
)

func (ap AdapterMySQL) DeleteFromTable(tableName string) string {
	return base.DeleteFromTable(tableName)
}
