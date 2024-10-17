package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client *mongo.Client
	db     *mongo.Database
}

// NewMongoClient connects to MongoDB and returns a MongoClient.
func NewMongoClient(mongoURL, dbName string) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	return &MongoClient{client: client, db: db}, nil
}
