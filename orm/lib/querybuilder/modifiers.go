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

// LeftJoinAndSelect adds relations to the query (eager loading)
func (qb *QueryBuilder) LeftJoinAndSelect(table string, condition string, selectFields ...string) *QueryBuilder {
	joinClause := JoinClause{
		JoinType:  "LEFT",
		Table:     table,
		Condition: condition,
		Select:    selectFields,
	}
	qb.joins = append(qb.joins, joinClause)
	return qb
}

func (qb *QueryBuilder) Debug() *QueryBuilder {
	qb.debug = true
	return qb
}
