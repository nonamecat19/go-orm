package querybuilder

import (
	"errors"
	"fmt"
)

// UpdateMany initializes an UPDATE query for the specified entity.
func (qb *QueryBuilder) UpdateMany(entity interface{}) error {
	tableName, _, _, err := qb.extractTableAndFields(entity)
	if err != nil {
		return err
	}

	if len(qb.where) == 0 {
		return errors.New("UPDATE query requires at least one WHERE clause to prevent accidental update of all rows")
	}

	if len(qb.set) == 0 {
		return errors.New("no updates provided")
	}

	query := fmt.Sprintf("UPDATE %s", tableName)

	query = qb.prepareSet(query)
	query = qb.prepareWhere(query)

	qb.query = query

	_, err = qb.ExecuteBuilderQuery()

	return err
}

func (qb *QueryBuilder) SetValues(values map[string]interface{}) *QueryBuilder {
	qb.set = values
	return qb
}
