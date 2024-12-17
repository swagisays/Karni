package karni

// FieldType defines possible data types
type FieldType string

const (
	String   FieldType = "string"
	Number   FieldType = "number"
	Boolean  FieldType = "boolean"
	ObjectID FieldType = "objectid"
	Date     FieldType = "date"
	Buffer   FieldType = "buffer"
)

// Field represents a schema field definition
type Field struct {
	Type       FieldType
	Required   bool
	Default    interface{}
	Unique     bool
	Trim       bool
	Lowercase  bool
	Validators []func(interface{}) error
}

// Schema represents a collection schema
type SchemaStruct struct {
	Fields map[string]Field
}

func Schema(definition map[string]Field) *SchemaStruct {
	return &SchemaStruct{Fields: definition}
}
