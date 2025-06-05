package adapter_postgres

import (
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
)

// Where adds a WHERE condition to the query
func (ap AdapterPostgres) Where(condition string, args []any) string {
	return base.Where(condition, args)
}

// AndWhere adds an AND WHERE to the query
func (ap AdapterPostgres) AndWhere(condition string, where string, args []any) string {
	return base.AndWhere(condition, where, args)
}

// OrWhere adds an OR WHERE to the query
func (ap AdapterPostgres) OrWhere(condition string, where string, args []any) string {
	return base.OrWhere(condition, where, args)
}
