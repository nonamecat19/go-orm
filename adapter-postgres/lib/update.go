package adapter_postgres

import (
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
)

func (ap AdapterPostgres) Update(tableName string) string {
	return base.Update(tableName)
}
