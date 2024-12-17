package karni

import (
	"github.com/swagisays/karni/karni/global"
	"go.mongodb.org/mongo-driver/mongo"
)

type ModelStruct struct {
	collection *mongo.Collection
	schema     *SchemaStruct
}

type Document struct {
	model *ModelStruct
}

func Model(collectionName string, schema *SchemaStruct) *ModelStruct {
	db := global.GetDB()
	collection := db.Collection(collectionName)

	return &ModelStruct{collection: collection, schema: schema}
}

func (m *ModelStruct) New() {

}
