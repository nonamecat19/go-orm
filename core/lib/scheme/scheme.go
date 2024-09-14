package scheme

type Field struct {
	Name        string
	Type        string
	Nullability bool
}

type TableScheme struct {
	Name   string
	Fields []Field
}
