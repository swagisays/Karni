package karni

import (
	"fmt"
	"time"

	"github.com/swagisays/karni/karni/global"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ModelStruct struct {
	collection *mongo.Collection
	schema     *SchemaStruct
}

func Model(collectionName string, schema *SchemaStruct) *ModelStruct {
	db := global.GetDB()
	fmt.Println(db)
	collection := db.Collection(collectionName)

	return &ModelStruct{collection: collection, schema: schema}
}

func (m *ModelStruct) New(data map[string]interface{}) Document {
	doc := map[string]interface{}{
		"_id":       primitive.NewObjectID(),
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
	}

	for key, value := range data {
		doc[key] = value
	}
	document := Document{model: m, data: doc}

	return document
}
