package adapter_postgres

import (
	base "adapter-base/lib"
)

func (ap AdapterPostgres) Update(tableName string) string {
	return base.Update(tableName)
}
