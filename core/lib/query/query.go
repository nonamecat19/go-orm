package query

type JoinClause struct {
	JoinType  string // "LEFT", "INNER", etc.
	Table     string
	Condition string
	Select    []string
}
