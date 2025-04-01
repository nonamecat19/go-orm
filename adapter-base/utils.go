package adapter_base

import (
	"fmt"
	"strings"
)

func JoinFields(fields []string) string {
	for i, field := range fields {
		if field == "<nil>" {
			fields[i] = "NULL"
		}
	}
	return strings.Join(fields, ", ")
}

func JoinFieldsStrictly(fields []string) string {
	mappedFields := make([]string, len(fields))
	for i, field := range fields {
		mappedFields[i] = fmt.Sprintf("%s AS \"%s\"", field, field)
	}
	return strings.Join(mappedFields, ", ")
}

// NormalizeSqlWithArgs change "?" to database valid syntax
func NormalizeSqlWithArgs(sql string, args []interface{}) string {
	placeholderIndex := len(args) + 1

	for {
		placeholder := fmt.Sprintf("$%d", placeholderIndex)
		sql = strings.Replace(sql, "?", placeholder, 1) // Replace only the first '?' occurrence
		if !strings.Contains(sql, "?") {
			break
		}
		placeholderIndex++
	}

	return sql
}
