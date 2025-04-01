package adapter_postgres

import base "adapter-base"

func (ap AdapterPostgres) DeleteFromTable(tableName string) string {
	return base.DeleteFromTable(tableName)
}
