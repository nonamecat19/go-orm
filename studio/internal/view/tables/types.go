package tables

type FieldInfo struct {
	Name string
	Type string
}

func GetFieldNames(fields []FieldInfo) []string {
	names := make([]string, len(fields))
	for i, field := range fields {
		names[i] = field.Name
	}
	return names
}
