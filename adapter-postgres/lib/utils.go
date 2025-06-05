package adapter_postgres

import (
	base "github.com/nonamecat19/go-orm/adapter-base/lib"
)

func (ap AdapterPostgres) JoinFieldsStrictly(fields []string) string {
	return base.JoinFieldsStrictly(fields)
}

func (ap AdapterPostgres) JoinFields(fields []string) string {
	return base.JoinFields(fields)
}

func (ap AdapterPostgres) NormalizeSqlWithArgs(sql string, args []any) string {
	return base.NormalizeSqlWithArgs(sql, args)
}
