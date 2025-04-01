package querybuilder

import (
	"errors"
)

// UpdateMany initializes an UPDATE query for the specified entity.
func (qb *QueryBuilder) UpdateMany(entity any) error {
	tableName, _, _, err := qb.extractTableAndFields(entity, true)
	if err != nil {
		return err
	}

	if len(qb.where) == 0 {
		return errors.New("UPDATE query requires at least one WHERE clause to prevent accidental update of all rows")
	}

	if len(qb.set) == 0 {
		return errors.New("no updates provided")
	}

	query := qb.adapter.Update(tableName)
	query = qb.prepareSet(query)
	query = qb.prepareWhere(query)

	qb.query = query

	_, err = qb.ExecuteBuilderQuery()

	return err
}

func (qb *QueryBuilder) SetValues(values map[string]any) *QueryBuilder {
	qb.set = values
	return qb
}
