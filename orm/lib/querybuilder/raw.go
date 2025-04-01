package querybuilder

import (
	"database/sql"
)

// ExecuteRaw executes raw sql with params
func (qb *QueryBuilder) ExecuteRaw(sql string, args ...any) (*sql.Rows, error) {
	qb.args = args
	qb.query = qb.adapter.NormalizeSqlWithArgs(sql, args)
	return qb.ExecuteBuilderQuery()
}
