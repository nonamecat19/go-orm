package querybuilder

import (
	"errors"
	"fmt"
)

// DeleteMany initializes a DELETE query for the specified entity.
func (qb *QueryBuilder) DeleteMany(entity interface{}) error {
	tableName, _, _, err := qb.extractTableAndFields(entity)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s", tableName)

	if len(qb.where) == 0 {
		return errors.New("DELETE query requires at least one WHERE clause to prevent accidental deletion of all rows")
	}

	query = qb.prepareWhere(query)

	qb.query = query

	_, err = qb.ExecuteQuery()

	return err
}
