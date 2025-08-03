package lib

import (
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	db   *mongo.Client
	once sync.Once
)

func InitMongoDB() error {
	var err error
	once.Do(func() {
		db, err = mongo.Connect(options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	})
	return err
}

func Db(coll string) *mongo.Collection {
	if db == nil {
		panic("mongo client is not initialized")
	}
	if coll == "" {
		panic("collection name must not be empty")
	}
	return db.Database("kiku").Collection(coll)
}
