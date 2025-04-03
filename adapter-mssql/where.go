package adapter_mssql

import (
	"fmt"
)

// Where adds a WHERE condition to the query
func (ap AdapterMSSQL) Where(condition string, args []any) string {
	newCondition := ap.NormalizeSqlWithArgs(condition, args)
	return fmt.Sprintf("(%s)", newCondition)
}

// AndWhere adds an AND WHERE to the query
func (ap AdapterMSSQL) AndWhere(condition string, where string, args []any) string {
	newCondition := ap.NormalizeSqlWithArgs(condition, args)
	return fmt.Sprintf("%s AND (%s)", where, newCondition)
}

// OrWhere adds an OR WHERE to the query
func (ap AdapterMSSQL) OrWhere(condition string, where string, args []any) string {
	newCondition := ap.NormalizeSqlWithArgs(condition, args)
	return fmt.Sprintf("%s OR (%s)", where, newCondition)
}
