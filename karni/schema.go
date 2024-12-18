package karni

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func GenerateValidator(schema *SchemaStruct) (bson.M, []mongo.IndexModel) {
	properties := bson.M{}
	requiredFields := []string{}
	indexModels := []mongo.IndexModel{}

	for fieldName, field := range schema.Fields {
		fieldSchema := bson.M{
			"bsonType": string(field.Type),
		}

		if field.Required {
			requiredFields = append(requiredFields, fieldName)
		}

		if field.Lowercase {
			fieldSchema["description"] = "This field must be lowercase."
		}

		if field.Unique {
			indexModel := mongo.IndexModel{
				Keys:    bson.D{{Key: fieldName, Value: 1}},
				Options: options.Index().SetUnique(true),
			}
			indexModels = append(indexModels, indexModel)
		}

		// Add more field-specific validations here if needed

		properties[fieldName] = fieldSchema
	}

	validator := bson.M{
		"$jsonSchema": bson.M{
			"bsonType":   "object",
			"required":   requiredFields,
			"properties": properties,
		},
	}

	return validator, indexModels
}
