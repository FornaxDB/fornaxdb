package schema

type SchemaValidator struct {

}

func NewSchemaValidator() SchemaValidator {
	return SchemaValidator{}
}

func (s ScalarFieldReturnType) Validate(schema *Schema) error {
	return nil
}

