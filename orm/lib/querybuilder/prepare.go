package querybuilder

//nolint:unused // used externally
//lint:ignore U1000 used externally
func (qb *QueryBuilder) prepareOrderBy(query string) string {
	return qb.adapter.PrepareOrderBy(query, qb.orderBy)
}

func (qb *QueryBuilder) prepareWhere(query string) string {
	return qb.adapter.PrepareWhere(query, qb.where)
}

//nolint:unused // used externally
//lint:ignore U1000 used externally
func (qb *QueryBuilder) prepareLimit(query string) string {
	return qb.adapter.PrepareLimit(query, qb.limit)
}

//nolint:unused // used externally
//lint:ignore U1000 used externally
func (qb *QueryBuilder) prepareOffset(query string) string {
	return qb.adapter.PrepareOffset(query, qb.offset)
}

func (qb *QueryBuilder) prepareSet(query string) string {
	normalizedQuery, newArgs := qb.adapter.PrepareSet(query, qb.set, qb.args)
	qb.args = newArgs

	return normalizedQuery
}
