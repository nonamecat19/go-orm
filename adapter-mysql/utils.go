package adapter_mysql

import base "adapter-base"

func (ap AdapterMySQL) JoinFieldsStrictly(fields []string) string {
	return base.JoinFieldsStrictly(fields)
}

func (ap AdapterMySQL) JoinFields(fields []string) string {
	return base.JoinFields(fields)
}

func (ap AdapterMySQL) NormalizeSqlWithArgs(sql string, args []any) string {
	return sql
}
