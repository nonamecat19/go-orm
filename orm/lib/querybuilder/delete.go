package querybuilder

import (
	"errors"
	"github.com/nonamecat19/go-orm/core/utils"
)

// DeleteMany initializes a DELETE query for the specified entity.
func (qb *QueryBuilder) DeleteMany(entity any) error {
	tableName, _, _, err := utils.ExtractTableAndFields(entity, true)
	if err != nil {
		return err
	}

	query := qb.adapter.DeleteFromTable(tableName)

	if len(qb.where) == 0 {
		return errors.New("DELETE query requires at least one WHERE clause to prevent accidental deletion of all rows")
	}

	query = qb.prepareWhere(query)

	qb.query = query

	_, err = qb.ExecuteBuilderQuery()

	return err
}
