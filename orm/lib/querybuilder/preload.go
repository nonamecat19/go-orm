package querybuilder

func (qb *QueryBuilder) Preload(column ...string) *QueryBuilder {
	qb.preloads = append(qb.preloads, column...)
	return qb
}
