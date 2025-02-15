package querybuilder

import (
	"errors"
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
}

// CreateQueryBuilder initializes a new QueryBuilder.
func CreateQueryBuilder(client client.DbClient) *QueryBuilder {
	return &QueryBuilder{
		client: client,
		limit:  -1, // Default to no limit
	}
}

// execute runs the built query.
func (qb *QueryBuilder) execute() error {
	if qb.query == "" {
		return errors.New("query not built; call FindMany first")
	}
	_, err := qb.client.Query(qb.query, qb.args...)
	return err
}
