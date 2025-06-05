package adapter_mysql

import (
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
)

func (ap AdapterMySQL) DeleteFromTable(tableName string) string {
	return base.DeleteFromTable(tableName)
}
