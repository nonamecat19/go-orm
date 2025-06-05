package adapter_mssql

import (
	base "adapter-base/lib"
	"fmt"
	"strings"
)

func (ap AdapterMSSQL) JoinFieldsStrictly(fields []string) string {
	return base.JoinFieldsStrictly(fields)
}

func (ap AdapterMSSQL) JoinFields(fields []string) string {
	return base.JoinFields(fields)
}

func (ap AdapterMSSQL) NormalizeSqlWithArgs(sql string, args []any) string {
	placeholderIndex := len(args) + 1

	for {
		placeholder := fmt.Sprintf("@p%d", placeholderIndex)
		sql = strings.Replace(sql, "?", placeholder, 1) // Replace only the first '?' occurrence
		if !strings.Contains(sql, "?") {
			break
		}
		placeholderIndex++
	}

	return sql
}
