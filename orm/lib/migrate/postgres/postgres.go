package postgres

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/lib/scheme"
	"github.com/nonamecat19/go-orm/core/utils"
)

var postgresMapping = map[string]string{
	"string": "varchar(256)",
	"int":    "integer",
	"int32":  "integer",
	"int64":  "integer",
	"float":  "numeric",
	"bool":   "boolean",
}

func GeneratePostgresTablesSQL(schemes []scheme.TableScheme) string {
	var sql string

	for _, item := range schemes {
		sql += fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", item.Name)

		for _, field := range item.Fields {
			fieldTypeStr := field.Type
			if fieldType, ok := postgresMapping[field.Type]; ok {
				fieldTypeStr = fieldType
			}

			sql += fmt.Sprintf("  %s %s%s,\n", field.Name, fieldTypeStr, utils.If(field.Nullability, " NULL", ""))
		}

		sql = sql[:len(sql)-2] + "\n);\n\n"
	}

	return sql
}
