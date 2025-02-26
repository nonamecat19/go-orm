package querybuilder

import "fmt"

// Where adds a WHERE condition to the query
func (qb *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
	newCondition := qb.normalizeCondition(condition)
	qb.where = fmt.Sprintf("(%s)", newCondition)
	qb.args = append(qb.args, args...)
	return qb
}

// AndWhere adds an AND WHERE to the query
func (qb *QueryBuilder) AndWhere(condition string, args ...interface{}) *QueryBuilder {
	newCondition := qb.normalizeCondition(condition)
	qb.where = fmt.Sprintf("%s AND (%s)", qb.where, newCondition)
	qb.args = append(qb.args, args...)
	return qb
}

// OrWhere adds an OR WHERE to the query
func (qb *QueryBuilder) OrWhere(condition string, args ...interface{}) *QueryBuilder {
	newCondition := qb.normalizeCondition(condition)
	qb.where = fmt.Sprintf("%s OR (%s)", qb.where, newCondition)
	qb.args = append(qb.args, args...)
	return qb
}
