package querybuilder

import (
	"database/sql"
)

// ExecuteRaw executes raw sql with params
func (qb *QueryBuilder) ExecuteRaw(sql string, args ...interface{}) (*sql.Rows, error) {
	normalizedSql := qb.normalizeSqlWithArgs(sql)
	qb.args = args
	qb.query = normalizedSql
	return qb.Query()
}
