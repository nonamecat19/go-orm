package adapter_postgres

import (
	base "adapter-base"
	"github.com/nonamecat19/go-orm/core/lib/query"
)

func (ap AdapterPostgres) PrepareOrderBy(query string, orderBy []string) string {
	return base.PrepareOrderBy(query, orderBy)
}

func (ap AdapterPostgres) PrepareWhere(query string, where string) string {
	return base.PrepareWhere(query, where)
}

func (ap AdapterPostgres) PrepareLimit(query string, limit int) string {
	return base.PrepareLimit(query, limit)
}

func (ap AdapterPostgres) PrepareOffset(query string, offset int) string {
	return base.PrepareOffset(query, offset)
}

func (ap AdapterPostgres) PrepareSet(query string, set map[string]any, args []any) (string, []any) {
	return base.PrepareSet(query, set, args)
}

func (ap AdapterPostgres) PrepareJoins(query string, joins []query.JoinClause) string {
	return base.PrepareJoins(query, joins)
}
