package schema

type DataValidator struct {}

func NewDataValidator() DataValidator {
	return DataValidator{}
}

func (d DataValidator) Validate(schema *Schema, data interface{}) {

}