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

func initGlobalContext() {
	if GlobalContext == nil {
		GlobalContext = &AppContext{}
	}
}
func SetClient(client *mongo.Client) {
	initGlobalContext()
	mu.Lock()
	defer mu.Unlock()
	GlobalContext.client = client
}

func SetDB(db *mongo.Database) {
	initGlobalContext()

	mu.Lock()
	defer mu.Unlock()
	GlobalContext.db = db
}

func GetClient() *mongo.Client {
	initGlobalContext()

	mu.Lock()
	defer mu.Unlock()
	return GlobalContext.client
}

func GetDB() *mongo.Database {
	initGlobalContext()

	mu.Lock()
	defer mu.Unlock()
	return GlobalContext.db
}

func SetGlobalContext(ctx *AppContext) {
	initGlobalContext()

	mu.Lock()
	defer mu.Unlock()
	GlobalContext = ctx
}

func GetGlobalContext() *AppContext {
	initGlobalContext()

	mu.Lock()
	defer mu.Unlock()
	return GlobalContext
}
