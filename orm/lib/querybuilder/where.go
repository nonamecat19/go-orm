package querybuilder

// Where adds a WHERE condition to the query
func (qb *QueryBuilder) Where(condition string, args ...any) *QueryBuilder {
	qb.where = qb.adapter.Where(condition, qb.args)
	qb.args = append(qb.args, args...)
	return qb
}

// AndWhere adds an AND WHERE to the query
func (qb *QueryBuilder) AndWhere(condition string, args ...any) *QueryBuilder {
	qb.where = qb.adapter.AndWhere(condition, qb.where, qb.args)
	qb.args = append(qb.args, args...)
	return qb
}

// OrWhere adds an OR WHERE to the query
func (qb *QueryBuilder) OrWhere(condition string, args ...any) *QueryBuilder {
	qb.where = qb.adapter.OrWhere(condition, qb.where, qb.args)
	qb.args = append(qb.args, args...)
	return qb
}
