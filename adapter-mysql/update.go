package adapter_mysql

import (
	base "adapter-base"
)

func (ap AdapterMySQL) Update(tableName string) string {
	return base.Update(tableName)
}
