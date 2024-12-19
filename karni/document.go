package karni

import (
	"context"
	"strings"

	karniErrors "github.com/swagisays/karni/karni/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Document struct {
	model *ModelStruct
	Data  map[string]interface{}
}

// Middleware function to process data
func ProcessData(data map[string]interface{}, schema *SchemaStruct) map[string]interface{} {
	for fieldName, field := range schema.Fields {
		if value, exists := data[fieldName]; exists {
			if field.Trim {
				if _, ok := value.(string); ok {
					data[fieldName] = strings.TrimSpace(data[fieldName].(string))

				}
			}
			if field.Lowercase {
				if _, ok := value.(string); ok {
					data[fieldName] = strings.ToLower(data[fieldName].(string))
				}
			}
			// Add more transformations as needed
		}
	}
	return data
}

func (d *Document) Save() (*mongo.InsertOneResult, error) {
	processedData := ProcessData(d.Data, d.model.schema)
	collection := d.model.collection

	result, err := collection.InsertOne(context.TODO(), bson.M(processedData))
	if err != nil {
		wrappedErr := karniErrors.HandleMongoValidationError(err)
		return nil, wrappedErr
	}
	return result, nil
}
