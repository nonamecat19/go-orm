package adapter_sqlite

import (
	base "adapter-base/lib"
)

func (ap AdapterSQLite) DeleteFromTable(tableName string) string {
	return base.DeleteFromTable(tableName)
}
