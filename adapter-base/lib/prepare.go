package adapter_base

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/lib/query"
)

func PrepareOrderBy(query string, orderBy []string) string {
	if len(orderBy) == 0 {
		return query
	}

	return query + " ORDER BY " + JoinFields(orderBy)
}

func PrepareWhere(query string, where string) string {
	if len(where) == 0 {
		return query
	}

	return fmt.Sprintf("%s WHERE %s", query, where)
}

func PrepareLimit(query string, limit int) string {
	if limit == -1 {
		return query
	}

	return query + fmt.Sprintf(" LIMIT %d", limit)
}

func PrepareOffset(query string, offset int) string {
	if offset == -1 {
		return query
	}

	return query + fmt.Sprintf(" OFFSET %d", offset)
}

func PrepareSet(query string, set map[string]any, args []any) (string, []any) {
	if len(set) == 0 {
		return query, args
	}

	newArgs := args

	var setClauses []string
	for column := range set {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", column))
	}

	queryWithSet := fmt.Sprintf("%s SET %s", query, JoinFields(setClauses))
	normalizedQuery := NormalizeSqlWithArgs(queryWithSet, args)
	for _, value := range set {
		newArgs = append(newArgs, value)
	}

	return normalizedQuery, newArgs
}

func PrepareJoins(query string, joins []query.JoinClause) string {
	newQuery := query
	for _, join := range joins {
		newQuery += fmt.Sprintf(" %s JOIN %s ON %s", join.JoinType, join.Table, join.Condition)
	}

	return newQuery
}
