package karni

import (
	"context"
	"fmt"
	"time"

	"github.com/swagisays/karni/karni/global"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ModelStruct struct {
	collection *mongo.Collection
	schema     *SchemaStruct
}

func Model(collectionName string, schema *SchemaStruct) *ModelStruct {
	db := global.GetDB()
	// Generate the validator using the schema
	validator, indexModels := GenerateValidator(schema)

	opts := options.CreateCollection().SetValidator(validator)
	err := db.CreateCollection(context.TODO(), collectionName, opts)
	if err != nil {
		// fmt.Println(err)

	}
	collection := db.Collection(collectionName)

	_, err = collection.Indexes().CreateMany(context.TODO(), indexModels)
	if err != nil {
		fmt.Println("Error creating index:", err)
	}

	return &ModelStruct{collection: collection, schema: schema}
}

func (m *ModelStruct) New(data map[string]interface{}) Document {
	// Process data using middleware
	processedData := ProcessData(data, m.schema)
	doc := map[string]interface{}{
		"_id":       primitive.NewObjectID(),
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
	}

	for key, value := range processedData {
		doc[key] = value
	}
	document := Document{model: m, Data: doc}

	return document
}
