package adapter_mysql

import (
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
)

func (ap AdapterMySQL) Update(tableName string) string {
	return base.Update(tableName)
}
