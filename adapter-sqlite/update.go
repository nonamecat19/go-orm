package adapter_sqlite

import (
	base "adapter-base"
)

func (ap AdapterSQLite) Update(tableName string) string {
	return base.Update(tableName)
}
