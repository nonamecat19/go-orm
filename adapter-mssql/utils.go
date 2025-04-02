package adapter_mssql

import base "adapter-base"

func (ap AdapterMSSQL) JoinFieldsStrictly(fields []string) string {
	return base.JoinFieldsStrictly(fields)
}

func (ap AdapterMSSQL) JoinFields(fields []string) string {
	return base.JoinFields(fields)
}

func (ap AdapterMSSQL) NormalizeSqlWithArgs(sql string, args []any) string {
	return base.NormalizeSqlWithArgs(sql, args)
}
