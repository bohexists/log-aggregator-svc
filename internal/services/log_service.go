package services

import (
	"log"

	"github.com/bohexists/log-aggregator-svc/internal/adapters/mongo"
	"github.com/bohexists/log-aggregator-svc/internal/domain"
)

// LogService orchestrates the flow of log entries from NATS to MongoDB.
type LogService struct {
	mongoRepo *mongo.MongoRepository
}

// NewLogService creates a new instance of LogService.
func NewLogService(mongoRepo *mongo.MongoRepository) *LogService {
	return &LogService{
		mongoRepo: mongoRepo,
	}
}

// ProcessLog обрабатывает полученный лог и сохраняет его в MongoDB.
func (s LogService) ProcessLog(logEntry domain.LogEntry) error {
	err := s.mongoRepo.InsertLog(&logEntry)
	if err != nil {
		log.Printf("Error saving log to MongoDB: %v", err)
		return err
	}
	log.Printf("Log successfully saved to MongoDB: %+v", logEntry)
	return nil
}
