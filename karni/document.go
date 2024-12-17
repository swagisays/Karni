package karni

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Document struct {
	model *ModelStruct
	data  map[string]interface{}
}

func (d *Document) Save() {
	collection := d.model.collection

	collection.InsertOne(context.TODO(), bson.M(d.data))
}
