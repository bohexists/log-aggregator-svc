package mongo

import (
	"context"
	"log"
	"time"

	"github.com/bohexists/log-aggregator-svc/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRepository provides an interface for interacting with the MongoDB database.
type MongoRepository struct {
	db *mongo.Database
}

// NewMongoRepository creates a new instance of MongoRepository.
func NewMongoRepository(client *mongo.Client, dbName string) *MongoRepository {
	return &MongoRepository{
		db: client.Database(dbName),
	}
}

// InsertLog saves a log entry in the MongoDB logs collection.
func (r *MongoRepository) InsertLog(logEntry *domain.LogEntry) error {
	collection := r.db.Collection("logs")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, logEntry)
	if err != nil {
		log.Printf("Error inserting log into MongoDB: %v", err)
		return err
	}
	log.Printf("Log entry inserted successfully: %+v", logEntry)
	return nil
}

// Close gracefully disconnects from MongoDB.
func (r *MongoRepository) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := r.db.Client().Disconnect(ctx)
	if err != nil {
		return err
	}
	log.Println("Disconnected from MongoDB")
	return nil
}
