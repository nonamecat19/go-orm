package querybuilder

// Select set selected columns from table
func (qb *QueryBuilder) Select(column ...string) *QueryBuilder {
	qb.selectFields = append(qb.selectFields, column...)
	return qb
}
