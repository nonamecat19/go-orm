package adapter_postgres

import base "adapter-base"

func (ap AdapterPostgres) JoinFieldsStrictly(fields []string) string {
	return base.JoinFieldsStrictly(fields)
}

func (ap AdapterPostgres) JoinFields(fields []string) string {
	return base.JoinFields(fields)
}

func (ap AdapterPostgres) NormalizeSqlWithArgs(sql string, args []interface{}) string {
	return base.NormalizeSqlWithArgs(sql, args)
}
