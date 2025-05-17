package querybuilder

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/nonamecat19/go-orm/core/lib/adapter"
	"github.com/nonamecat19/go-orm/core/lib/query"
	"github.com/nonamecat19/go-orm/orm/lib/client"
	"strings"
)

type QueryBuilder struct {
	client       client.DbClient
	adapter      adapter.Adapter
	query        string
	selectFields []string
	where        string
	orderBy      []string
	limit        int
	offset       int
	args         []any
	set          map[string]any
	joins        []query.JoinClause
	debug        bool
	preloads     []string
}

// CreateQueryBuilder initializes a new QueryBuilder.
func CreateQueryBuilder(client client.DbClient) *QueryBuilder {
	return &QueryBuilder{
		adapter: client.GetAdapter(),
		client:  client,
		limit:   -1, // Default to no limit
	}
}

func formatSQL(sql string) string {
	sql = strings.ReplaceAll(sql, "SELECT", "\nSELECT\n\t")
	sql = strings.ReplaceAll(sql, "FROM", "\nFROM")
	sql = strings.ReplaceAll(sql, "LEFT JOIN", "\nLEFT JOIN")
	sql = strings.ReplaceAll(sql, "WHERE", "\nWHERE")
	sql = strings.ReplaceAll(sql, "AND", "\n\tAND")
	sql = strings.ReplaceAll(sql, "LIMIT", "\nLIMIT")
	sql = strings.ReplaceAll(sql, "TOP", "\nTOP")
	sql = strings.ReplaceAll(sql, "OFFSET", "\nOFFSET")
	sql = strings.ReplaceAll(sql, "ORDER BY", "\nORDER BY")

	sql = strings.ReplaceAll(sql, "ON", "\n\tON")

	return sql
}

// ExecuteBuilderQuery runs the built query with builder options.
func (qb *QueryBuilder) ExecuteBuilderQuery() (*sql.Rows, error) {
	return qb.ExecuteQuery(qb.query, qb.args...)
}

// ExecuteQuery runs the query.
func (qb *QueryBuilder) ExecuteQuery(query string, args ...any) (*sql.Rows, error) {
	if query == "" {
		return nil, errors.New("query not built")
	}

	preparedQuery, preparedArgs := qb.adapter.PrepareQueryAndArgs(query, args)
	if qb.debug {
		// Green color for the query
		fmt.Println("\033[32m" + formatSQL(preparedQuery) + "\033[0m")
		// Yellow color for args
		fmt.Printf("\033[33m%v\033[0m\n", args)
	}

	return qb.client.GetDb().Query(preparedQuery, preparedArgs...)
}
