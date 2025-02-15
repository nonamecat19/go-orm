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

	return query + " WHERE " + strings.Join(qb.where, ", ")
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
