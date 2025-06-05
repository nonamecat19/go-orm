package adapter_mysql

import (
	"fmt"
)

// Where adds a WHERE condition to the query
func (ap AdapterMySQL) Where(condition string, args []any) string {
	newCondition := ap.NormalizeSqlWithArgs(condition, args)
	return fmt.Sprintf("(%s)", newCondition)
}

// AndWhere adds an AND WHERE to the query
func (ap AdapterMySQL) AndWhere(condition string, where string, args []any) string {
	newCondition := ap.NormalizeSqlWithArgs(condition, args)
	return fmt.Sprintf("%s AND (%s)", where, newCondition)
}

// OrWhere adds an OR WHERE to the query
func (ap AdapterMySQL) OrWhere(condition string, where string, args []any) string {
	newCondition := ap.NormalizeSqlWithArgs(condition, args)
	return fmt.Sprintf("%s OR (%s)", where, newCondition)
}
