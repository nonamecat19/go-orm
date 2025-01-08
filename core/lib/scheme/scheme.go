package scheme

type Field struct {
	Name        string
	Type        string
	Nullability bool
	Unique      bool
	PrimaryKey  bool
}

type TableScheme struct {
	Name   string
	Fields []Field
}
