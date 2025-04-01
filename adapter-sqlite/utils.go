package adapter_sqlite

import base "adapter-base"

func (ap AdapterSQLite) JoinFieldsStrictly(fields []string) string {
	return base.JoinFieldsStrictly(fields)
}

func (ap AdapterSQLite) JoinFields(fields []string) string {
	return base.JoinFields(fields)
}

func (ap AdapterSQLite) NormalizeSqlWithArgs(sql string, args []any) string {
	return base.NormalizeSqlWithArgs(sql, args)
}
