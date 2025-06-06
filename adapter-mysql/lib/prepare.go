package adapter_mysql

import (
	"fmt"
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
	"github.com/nonamecat19/go-orm/core/lib/query"
)

func (ap AdapterMySQL) PrepareOrderBy(query string, orderBy []string) string {
	return base.PrepareOrderBy(query, orderBy)
}

func (ap AdapterMySQL) PrepareWhere(query string, where string) string {
	return base.PrepareWhere(query, where)
}

func (ap AdapterMySQL) PrepareLimit(query string, limit int) string {
	return base.PrepareLimit(query, limit)
}

func (ap AdapterMySQL) PrepareOffset(query string, offset int) string {
	return base.PrepareOffset(query, offset)
}

func (ap AdapterMySQL) PrepareSet(query string, set map[string]any, args []any) (string, []any) {
	if len(set) == 0 {
		return query, args
	}

	newArgs := args

	var setClauses []string
	for column := range set {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", column))
	}

	queryWithSet := fmt.Sprintf("%s SET %s", query, ap.JoinFields(setClauses))
	normalizedQuery := ap.NormalizeSqlWithArgs(queryWithSet, args)
	for _, value := range set {
		newArgs = append(newArgs, value)
	}

	return normalizedQuery, newArgs
}

func (ap AdapterMySQL) PrepareJoins(query string, joins []query.JoinClause) string {
	return base.PrepareJoins(query, joins)
}

func (ap AdapterMySQL) PrepareQueryAndArgs(query string, args []any) (string, []any) {
	return query, args
}
