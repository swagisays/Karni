package karni

import (
	"context"
	"time"

	liberrors "github.com/swagisays/karni/karni/errors"
	"github.com/swagisays/karni/karni/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	global.SetClient(client)
	return client, nil
}

func Connect(url string, name string) error {
	client, err := createClient(url)
	if err != nil {
		return liberrors.WrapError(1002, "Failed to create MongoDB client", err)
	}
	db := client.Database(name)
	global.SetDB(db)
	return nil
}
