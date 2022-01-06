package schema


type FieldType string
const (
	STRING_FIELD_TYPE = "string"
	INT_FIELD_TYPE    = "int"
	BOOL_FIELD_TYPE   = "bool"
	FLOAT_FIELD_TYPE  = "float"

)

type Field struct {
	Name string
	Type FieldType
	Operators string
}

type Type struct {
	Name string
	Fields []Field
}

type Relation struct {
	Name string
	Fields []Field
	Src Type
	Des Type
}

type Schema struct {
	Types []Type
	Relation []Relation
}