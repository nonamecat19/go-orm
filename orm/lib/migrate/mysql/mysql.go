package mysql

var mysqlMapping = map[string]string{
	"string": "varchar(255)",
	"int":    "integer",
	"int32":  "integer",
	"int64":  "bigint",
	"float":  "double precision",
	"bool":   "boolean",
}
