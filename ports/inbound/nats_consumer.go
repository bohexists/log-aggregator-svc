package inbound

import (
	"github.com/bohexists/log-aggregator-svc/internal/adapters/mongo"
	"github.com/bohexists/log-aggregator-svc/internal/domain"
	"log"
)

// NatsConsumer is responsible for processing log entries received from NATS and saving them to MongoDB.
type NatsConsumer struct {
	mongoRepository mongo.MongoRepository // Dependency Injection for Mongo repository interface
}

// NewNatsConsumer creates a new instance of NatsConsumer with a Mongo repository dependency.
func NewNatsConsumer(mongoRepository mongo.MongoRepository) *NatsConsumer {
	return &NatsConsumer{
		mongoRepository: mongoRepository,
	}
}

// ConsumeLog processes the log entry and saves it to MongoDB using the repository.
func (c *NatsConsumer) ConsumeLog(logEntry *domain.LogEntry) error {

	err := c.mongoRepository.InsertLog(logEntry)
	if err != nil {
		log.Printf("Error saving log to MongoDB: %v", err)
		return err
	}

	log.Printf("Log successfully saved to MongoDB: %+v", logEntry)
	return nil
}
