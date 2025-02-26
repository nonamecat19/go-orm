package querybuilder

// OrderBy adds order by clauses to the query
func (qb *QueryBuilder) OrderBy(fields ...string) *QueryBuilder {
	qb.orderBy = append(qb.orderBy, fields...)
	return qb
}

// Limit sets a limit for the query results
func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.limit = limit
	return qb
}

// Offset sets an offset for the query results
func (qb *QueryBuilder) Offset(offset int) *QueryBuilder {
	qb.offset = offset
	return qb
}

// WithRelations adds relations to the query (eager loading)
func (qb *QueryBuilder) WithRelations(relations ...string) *QueryBuilder {
	qb.relations = append(qb.relations, relations...)
	return qb
}

func (qb *QueryBuilder) Debug() *QueryBuilder {
	qb.debug = true
	return qb
}
