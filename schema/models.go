package schema

type Schema struct {
	Types     map[string]Type
	Relations map[string]Relation
}

type Type struct {
	Name   string
	Fields []Field
}

type Relation struct {
	Name        string
	Fields      []Field
	Source      *Type
	Destination *Type
}

type Field struct {
	Key        string
	Arguments  []Argument
	ReturnType FieldReturnType
}

type Argument struct {
	Key         string
	RetunType    FieldReturnType
	DefaultValue interface{}
}

type FieldReturnType interface {
	isFieldReturnType()
}

type ScalarFieldReturnType struct {
	Type     string
	Nullable bool
}

func (ScalarFieldReturnType) isFieldReturnType() {}

type VectorFieldReturnType struct {
	Type              string
	IsElementNullable bool
	IsVectorNullable  bool
}

func (VectorFieldReturnType) isFieldReturnType() {}

func NewSchema() *Schema {
	return &Schema{
		Types: make(map[string]Type),
		Relations: make(map[string]Relation),
	}
}