package adapter_mssql

import (
	base "adapter-base"
	"fmt"
	"github.com/nonamecat19/go-orm/core/lib/query"
)

func (ap AdapterMSSQL) PrepareOrderBy(query string, orderBy []string) string {
	return base.PrepareOrderBy(query, orderBy)
}

func (ap AdapterMSSQL) PrepareWhere(query string, where string) string {
	return base.PrepareWhere(query, where)
}

func (ap AdapterMSSQL) PrepareLimit(query string, limit int) string {
	if limit == -1 {
		return query
	}

	return query + fmt.Sprintf(" FETCH NEXT %d ROWS ONLY", limit)
}

func (ap AdapterMSSQL) PrepareOffset(query string, offset int) string {
	if offset == -1 {
		return query
	}

	return query + fmt.Sprintf(" OFFSET %d ROWS", offset)
}

func (ap AdapterMSSQL) PrepareSet(query string, set map[string]any, args []any) (string, []any) {
	return base.PrepareSet(query, set, args)
}

func (ap AdapterMSSQL) PrepareJoins(query string, joins []query.JoinClause) string {
	return base.PrepareJoins(query, joins)
}
