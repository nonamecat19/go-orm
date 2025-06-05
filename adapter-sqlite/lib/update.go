package adapter_sqlite

import (
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
)

func (ap AdapterSQLite) Update(tableName string) string {
	return base.Update(tableName)
}
