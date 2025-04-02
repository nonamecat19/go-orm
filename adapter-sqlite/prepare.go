package adapter_sqlite

import (
	base "adapter-base"
	"github.com/nonamecat19/go-orm/core/lib/query"
)

func (ap AdapterSQLite) PrepareOrderBy(query string, orderBy []string) string {
	return base.PrepareOrderBy(query, orderBy)
}

func (ap AdapterSQLite) PrepareWhere(query string, where string) string {
	return base.PrepareWhere(query, where)
}

func (ap AdapterSQLite) PrepareLimit(query string, limit int) string {
	return base.PrepareLimit(query, limit)
}

func (ap AdapterSQLite) PrepareOffset(query string, offset int) string {
	return base.PrepareOffset(query, offset)
}

func (ap AdapterSQLite) PrepareSet(query string, set map[string]any, args []any) (string, []any) {
	return base.PrepareSet(query, set, args)
}

func (ap AdapterSQLite) PrepareJoins(query string, joins []query.JoinClause) string {
	return base.PrepareJoins(query, joins)
}
