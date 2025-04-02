package adapter_mssql

import (
	base "adapter-base"
)

// Where adds a WHERE condition to the query
func (ap AdapterMSSQL) Where(condition string, args []any) string {
	return base.Where(condition, args)
}

// AndWhere adds an AND WHERE to the query
func (ap AdapterMSSQL) AndWhere(condition string, where string, args []any) string {
	return base.AndWhere(condition, where, args)
}

// OrWhere adds an OR WHERE to the query
func (ap AdapterMSSQL) OrWhere(condition string, where string, args []any) string {
	return base.OrWhere(condition, where, args)
}
