package querybuilder

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/nonamecat19/go-orm/orm/lib/client"
)

type QueryBuilder struct {
	client    client.DbClient
	query     string
	where     []string
	orderBy   []string
	limit     int
	offset    int
	relations []string
	args      []interface{}
	debug     bool
}

// CreateQueryBuilder initializes a new QueryBuilder.
func CreateQueryBuilder(client client.DbClient) *QueryBuilder {
	return &QueryBuilder{
		client: client,
		limit:  -1, // Default to no limit
	}
}

// Query runs the built query.
func (qb *QueryBuilder) Query() (*sql.Rows, error) {
	if qb.debug {
		fmt.Println(qb.query, qb.args)
	}
	if qb.query == "" {
		return nil, errors.New("query not built")
	}

	return qb.client.GetDb().Query(qb.query, qb.args...)
}
