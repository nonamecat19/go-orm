package querybuilder

import (
	"errors"
	"fmt"
	"github.com/nonamecat19/go-orm/orm/lib/client"
	"reflect"
	"strings"
)

type QueryBuilder struct {
	client    client.DbClient
	query     string
	where     []string
	orderBy   []string
	limit     int
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

// FindMany initializes a SELECT query for the specified entity.
func (qb *QueryBuilder) FindMany(entity interface{}) (*string, error) {
	entityType := reflect.TypeOf(entity)
	if entityType.Kind() != reflect.Ptr || entityType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("entity must be a pointer to a struct")
	}

	entityType = entityType.Elem()
	tableName := ""
	if tableNameMethod, ok := reflect.New(entityType).Interface().(interface{ TableName() string }); ok {
		tableName = tableNameMethod.TableName()
	} else {
		return nil, errors.New("entity struct must implement TableName() string method")
	}

	var fieldNames []string
	for i := 0; i < entityType.NumField(); i++ {
		fieldTag := entityType.Field(i).Tag.Get("db")
		if fieldTag != "" {
			fieldNames = append(fieldNames, fieldTag)
		}
	}

	query := fmt.Sprintf("SELECT %s FROM %s", joinFields(fieldNames), tableName)

	if len(qb.where) > 0 {
		query += " WHERE " + strings.Join(qb.where, " AND ")
	}

	if len(qb.orderBy) > 0 {
		query += " ORDER BY " + strings.Join(qb.orderBy, ", ")
	}

	if qb.limit > -1 {
		query += fmt.Sprintf(" LIMIT %d", qb.limit)
	}

	qb.query = query
	return &query, nil
}

// Where adds a where condition to the query.
func (qb *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
	qb.where = append(qb.where, condition)
	qb.args = append(qb.args, args...)
	return qb
}

// OrderBy adds order by clauses to the query.
func (qb *QueryBuilder) OrderBy(fields ...string) *QueryBuilder {
	qb.orderBy = append(qb.orderBy, fields...)
	return qb
}

// Limit sets a limit for the query results.
func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.limit = limit
	return qb
}

// WithRelations adds relations to the query (eager loading).
func (qb *QueryBuilder) WithRelations(relations ...string) *QueryBuilder {
	qb.relations = append(qb.relations, relations...)
	return qb
}

// execute runs the built query.
func (qb *QueryBuilder) execute() error {
	if qb.query == "" {
		return errors.New("query not built; call FindMany first")
	}
	_, err := qb.client.Query(qb.query, qb.args...)
	return err
}

// UpdateMany initializes an UPDATE query for the specified entity.
func (qb *QueryBuilder) UpdateMany(entity interface{}, updates map[string]interface{}) (*string, error) {
	tableName, _, err := qb.extractTableAndFields(entity)
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

	if len(qb.where) > 0 {
		query += " WHERE " + strings.Join(qb.where, " AND ")
	}

	qb.query = query
	return &query, nil
}

// Delete initializes a DELETE query for the specified entity.
func (qb *QueryBuilder) Delete(entity interface{}) (*string, error) {
	tableName, _, err := qb.extractTableAndFields(entity)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("DELETE FROM %s", tableName)

	if len(qb.where) > 0 {
		query += " WHERE " + strings.Join(qb.where, " AND ")
	} else {
		return nil, errors.New("DELETE query requires at least one WHERE clause to prevent accidental deletion of all rows")
	}

	qb.query = query
	return &query, nil
}

// extractTableAndFields: Extracts table name and field names from an entity.
func (qb *QueryBuilder) extractTableAndFields(entity interface{}) (string, []string, error) {
	entityType := reflect.TypeOf(entity)
	if entityType.Kind() != reflect.Ptr || entityType.Elem().Kind() != reflect.Struct {
		return "", nil, errors.New("entity must be a pointer to a struct")
	}

	entityType = entityType.Elem()
	tableName := ""
	if tableNameMethod, ok := reflect.New(entityType).Interface().(interface{ TableName() string }); ok {
		tableName = tableNameMethod.TableName()
	} else {
		return "", nil, errors.New("entity struct must implement TableName() string method")
	}

	var fieldNames []string
	for i := 0; i < entityType.NumField(); i++ {
		fieldTag := entityType.Field(i).Tag.Get("db")
		if fieldTag != "" {
			fieldNames = append(fieldNames, fieldTag)
		}
	}

	return tableName, fieldNames, nil
}

func joinFields(fields []string) string {
	return strings.Join(fields, ", ")
}
