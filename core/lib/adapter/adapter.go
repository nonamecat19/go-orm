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
	PrepareSet(query string, set map[string]interface{}, args []interface{}) (string, []interface{})
	Where(condition string, where string, args ...interface{}) string
	AndWhere(condition string, where string, args ...interface{}) string
	OrWhere(condition string, where string, args ...interface{}) string
	Insert(tableName string, fieldNames []string, values []any, args []any) (string, []any)
}
