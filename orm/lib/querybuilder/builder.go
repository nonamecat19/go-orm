package querybuilder

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/nonamecat19/go-orm/orm/lib/client"
	"strings"
)

type JoinClause struct {
	JoinType  string // "LEFT", "INNER", etc.
	Table     string
	Condition string
	Select    []string
}

type QueryBuilder struct {
	client       client.DbClient
	query        string
	selectFields []string
	where        string
	orderBy      []string
	limit        int
	offset       int
	relations    []string
	args         []interface{}
	joins        []JoinClause
	debug        bool
}

// CreateQueryBuilder initializes a new QueryBuilder.
func CreateQueryBuilder(client client.DbClient) *QueryBuilder {
	return &QueryBuilder{
		client: client,
		limit:  -1, // Default to no limit
	}
}

func formatSQL(sql string) string {
	sql = strings.ReplaceAll(sql, "SELECT", "\nSELECT\n\t")
	sql = strings.ReplaceAll(sql, "FROM", "\nFROM")
	sql = strings.ReplaceAll(sql, "LEFT JOIN", "\nLEFT JOIN")
	sql = strings.ReplaceAll(sql, "WHERE", "\nWHERE")
	sql = strings.ReplaceAll(sql, "AND", "\n\tAND")
	sql = strings.ReplaceAll(sql, "LIMIT", "\nLIMIT")
	sql = strings.ReplaceAll(sql, "OFFSET", "\nOFFSET")

	sql = strings.ReplaceAll(sql, "ON", "\n\tON")

	return sql
}

// ExecuteQuery runs the built query.
func (qb *QueryBuilder) ExecuteQuery() (*sql.Rows, error) {
	if qb.debug {
		// Green color for the query
		fmt.Println("\033[32m" + formatSQL(qb.query) + "\033[0m")
		// Yellow color for args
		fmt.Printf("\033[33m%v\033[0m\n", qb.args)
	}
	if qb.query == "" {
		return nil, errors.New("query not built")
	}

	return qb.client.GetDb().Query(qb.query, qb.args...)
}
