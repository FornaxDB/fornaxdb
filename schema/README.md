# Schema for FornaxDB

## Grammar for Schema

```
schema -> ("type" identifier "{" schema_def  "}") | 
					("relation" identifier "{" relationship_def "}") ;

identifier ->  (a-z | A-Z | 0-9 | _ )* ;

schema_def -> (identifier "(" argument*  ")" ":" return_type ",")* ;

argument -> indentifer ":" return_type ("= " identifier | "") 

return_type -> (scalar_return_type | vector_return_type) ;

scalar_return_type -> ("string" | "int" | "float" | "boolean")("!" | "?" | "") ;

vector_return_type -> "[" scalar_return_type "]" ("!" | "?" | "") ;
```