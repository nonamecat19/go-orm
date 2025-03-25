package querybuilder

import (
	"fmt"
	"strings"
)

func (qb *QueryBuilder) prepareOrderBy(query string) string {
	if len(qb.orderBy) == 0 {
		return query
	}

	return query + " ORDER BY " + strings.Join(qb.orderBy, ", ")
}

func (qb *QueryBuilder) prepareWhere(query string) string {
	if len(qb.where) == 0 {
		return query
	}

	return fmt.Sprintf("%s WHERE %s", query, qb.where)
}

func (qb *QueryBuilder) prepareLimit(query string) string {
	if qb.limit == -1 {
		return query
	}

	return query + fmt.Sprintf(" LIMIT %d", qb.limit)
}

func (qb *QueryBuilder) prepareOffset(query string) string {
	if qb.limit == -1 {
		return query
	}

	return query + fmt.Sprintf(" OFFSET %d", qb.offset)
}

func (qb *QueryBuilder) prepareSet(query string) string {
	if len(qb.set) == 0 {
		return query
	}

	var setClauses []string
	for column := range qb.set {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", column))
	}

	queryWithSet := fmt.Sprintf("%s SET %s", query, JoinFields(setClauses))
	normalizedQuery := qb.normalizeSqlWithArgs(queryWithSet)
	for _, value := range qb.set {
		qb.args = append(qb.args, value)
	}

	return normalizedQuery
}
