package sqlite

var sqliteMapping = map[string]string{
	"string": "text",
	"int":    "integer",
	"int32":  "integer",
	"int64":  "integer",
	"float":  "real",
	"bool":   "integer",
}
