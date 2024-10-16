package mongo

import (
	"context"
	"log"
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

// InsertLog inserts a log entry into the logs collection.
func (m *MongoClient) InsertLog(collection string, document interface{}) error {
	_, err := m.db.Collection(collection).InsertOne(context.Background(), document)
	if err != nil {
		log.Printf("Error inserting log: %v", err)
		return err
	}
	return nil
}
