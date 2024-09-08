package infrastructure

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MyMongoDB struct{
	Client *mongo.Client
}

func (db *MyMongoDB) getUri() (uri string) {
	username := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	uri = fmt.Sprintf("mongodb://%s:%s@mongo:27017",username,password)
	return uri
}

func (db *MyMongoDB) Connect(ctx context.Context) (err error) {
	opt := options.Client().ApplyURI(db.getUri())
	if err := opt.Validate(); err != nil{
		return err
	}

	db.Client, err = mongo.Connect(ctx, opt)
	return err

}

func (db *MyMongoDB) Disconnect(ctx context.Context) error {
	return db.Client.Disconnect(ctx)	
}

func (db *MyMongoDB) Ping(ctx context.Context) error {
	if err := db.Client.Ping(ctx, nil); err != nil{
		return err
	}
	fmt.Println("Successfully connect")
	return nil
}