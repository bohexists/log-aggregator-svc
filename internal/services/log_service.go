package service

import (
	"github.com/bohexists/log-aggregator-svc/internal/adapters/mongo"
	"github.com/bohexists/log-aggregator-svc/internal/domain"
	"github.com/bohexists/log-aggregator-svc/ports/inbound"
	"log"
)

// LogService orchestrates the flow of log entries from NATS to MongoDB.
type LogService struct {
	natsConsumer *inbound.NatsConsumer
	mongoRepo    *mongo.MongoRepository
}

// NewLogService creates a new instance of LogService.
func NewLogService(natsConsumer *inbound.NatsConsumer, mongoRepo *mongo.MongoRepository) *LogService {
	return &LogService{
		natsConsumer: natsConsumer,
		mongoRepo:    mongoRepo,
	}
}

// Start begins the subscription process and listens for logs to process.
func (s *LogService) Start(subject string) error {
	err := s.natsConsumer.ConsumeLog(subject)
	if err != nil {
		log.Printf("Error subscribing to NATS: %v", err)
		return err
	}

	log.Printf("Successfully subscribed to NATS subject: %s", subject)
	return nil
}

// SaveLog saves a log entry to MongoDB.
func (s *LogService) SaveLog(logEntry domain.LogEntry) error {
	err := s.mongoRepo.InsertLog(logEntry)
	if err != nil {
		log.Printf("Error saving log to MongoDB: %v", err)
		return err
	}
	log.Printf("Log successfully saved to MongoDB: %+v", logEntry)
	return nil
}
