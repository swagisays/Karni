package global

import (
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext struct {
	client *mongo.Client
	db     *mongo.Database
}

var (
	GlobalContext *AppContext
	mu            sync.Mutex
)

func SetClient(client *mongo.Client) {
	mu.Lock()
	defer mu.Unlock()
	GlobalContext.client = client
}

func SetDB(db *mongo.Database) {
	mu.Lock()
	defer mu.Unlock()
	GlobalContext.db = db
}

func GetClient() *mongo.Client {
	mu.Lock()
	defer mu.Unlock()
	return GlobalContext.client
}

func GetDB() *mongo.Database {
	mu.Lock()
	defer mu.Unlock()
	return GlobalContext.db
}

func SetGlobalContext(ctx *AppContext) {
	mu.Lock()
	defer mu.Unlock()
	GlobalContext = ctx
}

func GetGlobalContext() *AppContext {
	mu.Lock()
	defer mu.Unlock()
	return GlobalContext
}
