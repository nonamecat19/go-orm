package querybuilder

import (
	"fmt"
)

// FindOne initialized a SELECT query for one record of the specified entity
func (qb *QueryBuilder) FindOne() {
}

// FindMany initializes a SELECT query for the specified entity.
func (qb *QueryBuilder) FindMany(entity interface{}) error {
	tableName, fieldNames, err := qb.extractTableAndFields(entity)

	if err != nil {
		return err
	}

	query := fmt.Sprintf("SELECT %s FROM %s", joinFields(fieldNames), tableName)

	query = qb.prepareWhere(query)
	query = qb.prepareOrderBy(query)
	query = qb.prepareLimit(query)
	query = qb.prepareOffset(query)

	qb.query = query

	println(qb.query)
	println(qb.args)
	a, err := qb.client.GetDb().Query(qb.query, qb.args...)
	println(a)

	return err
}
