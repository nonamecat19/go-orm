package adapter

import (
	"github.com/nonamecat19/go-orm/core/lib/config"
	"github.com/nonamecat19/go-orm/core/lib/query"
)

type DbAdapter interface {
	GetConnString(config config.ORMConfig) string
	GetDbDriver() string
	DeleteFromTable(tableName string) string
	PrepareOrderBy(query string, orderBy []string) string
	PrepareWhere(query string, where string) string
	PrepareLimit(query string, limit int) string
	PrepareOffset(query string, offset int) string
	PrepareSet(query string, set map[string]any, args []any) (string, []any)
	PrepareJoins(query string, joins []query.JoinClause) string
	Where(condition string, args []any) string
	AndWhere(condition string, where string, args []any) string
	OrWhere(condition string, where string, args []any) string
	Insert(tableName string, fieldNames []string, values []any, args []any) (string, []any)
	JoinFields(fields []string) string
	JoinFieldsStrictly(fields []string) string
	NormalizeSqlWithArgs(sql string, args []any) string
	Update(tableName string) string
	GetReadQuery(tableName string, fields []string, fromSubquery string) string
	GetFromSubquery(tableName string, where string, orderBy []string, limit int, offset int) string
	GetSelectQuery(selectValue string, fromValue string) string
	GetSelectWhereIn(tableName string, selectValue string, fieldName string, fieldValues []string) string
	PrepareQueryAndArgs(query string, args []any) (string, []any)
}
