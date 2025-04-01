package adapter

import "github.com/nonamecat19/go-orm/core/lib/config"

type Adapter interface {
	GetConnString(config config.ORMConfig) string
	GetDbDriver() string
	DeleteFromTable(tableName string) string
	PrepareOrderBy(query string, orderBy []string) string
	PrepareWhere(query string, where string) string
	PrepareLimit(query string, limit int) string
	PrepareOffset(query string, offset int) string
	PrepareSet(query string, set map[string]any, args []any) (string, []any)
	Where(condition string, where string, args ...any) string
	AndWhere(condition string, where string, args ...any) string
	OrWhere(condition string, where string, args ...any) string
	Insert(tableName string, fieldNames []string, values []any, args []any) (string, []any)
	NormalizeSqlWithArgs(sql string, args []any) string
}
