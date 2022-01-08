package errors

import (
	"github.com/joomcode/errorx"
)

var (
	ParserErrors = errorx.NewNamespace("parser")
	SchemaErrors = errorx.NewNamespace("schema")
	QueryErrors  = errorx.NewNamespace("query")
	StorageErrors = errorx.NewNamespace("storage")
)

var ( 
	InvalidToken = ParserErrors.NewType("invalid_token")	
)

var (
	SchemaInvalid = SchemaErrors.NewType("invalid")
	SchemaNotFound = SchemaErrors.NewType("not_found")
	SchemaAlreadyExists = SchemaErrors.NewType("already_exists")
	SchemaInvalidName = SchemaErrors.NewType("invalid_name")
	SchemaInvalidType = SchemaErrors.NewType("invalid_type")
	SchemaInvalidField = SchemaErrors.NewType("invalid_field")
	SchemaInvalidOperator = SchemaErrors.NewType("invalid_operator")
)
