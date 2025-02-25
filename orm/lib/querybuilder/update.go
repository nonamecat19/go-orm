package querybuilder

import (
	"errors"
	"fmt"
	"strings"
)

// UpdateMany initializes an UPDATE query for the specified entity.
func (qb *QueryBuilder) UpdateMany(entity interface{}, updates map[string]interface{}) (*string, error) {
	tableName, _, _, err := qb.extractTableAndFields(entity)
	if err != nil {
		return nil, err
	}

	if len(updates) == 0 {
		return nil, errors.New("no updates provided")
	}

	var setClauses []string
	for column, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", column))
		qb.args = append(qb.args, value)
	}

	query := fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(setClauses, ", "))

	query = qb.prepareWhere(query)

	qb.query = query
	return &query, nil
}
