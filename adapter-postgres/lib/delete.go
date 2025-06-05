package adapter_postgres

import (
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
)

func (ap AdapterPostgres) DeleteFromTable(tableName string) string {
	return base.DeleteFromTable(tableName)
}
