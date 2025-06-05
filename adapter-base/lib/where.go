package adapter_base

import "fmt"

// Where adds a WHERE condition to the query
func Where(condition string, args []any) string {
	newCondition := NormalizeSqlWithArgs(condition, args)
	return fmt.Sprintf("(%s)", newCondition)
}

// AndWhere adds an AND WHERE to the query
func AndWhere(condition string, where string, args []any) string {
	newCondition := NormalizeSqlWithArgs(condition, args)
	return fmt.Sprintf("%s AND (%s)", where, newCondition)
}

// OrWhere adds an OR WHERE to the query
func OrWhere(condition string, where string, args []any) string {
	newCondition := NormalizeSqlWithArgs(condition, args)
	return fmt.Sprintf("%s OR (%s)", where, newCondition)
}
