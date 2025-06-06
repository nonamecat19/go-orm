package adapter_mssql

import (
	"database/sql"
	"fmt"
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
	"github.com/nonamecat19/go-orm/core/lib/query"
	"github.com/nonamecat19/go-orm/core/utils"
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

func (ap AdapterMSSQL) PrepareJoins(query string, joins []query.JoinClause) string {
	return base.PrepareJoins(query, joins)
}

func (ap AdapterMSSQL) PrepareQueryAndArgs(query string, args []any) (string, []any) {
	return query, utils.MapWithIndex(args, func(arg any, index int) any {
		return sql.Named(fmt.Sprintf("p%d", index+1), arg)
	})
}
