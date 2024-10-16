package nats

import (
	"encoding/json"
	"log"

	"github.com/bohexists/log-aggregator-svc/internal/adapters/mongo"
	"github.com/bohexists/log-aggregator-svc/internal/domain"
	"github.com/nats-io/nats.go"
)

// NatsSubscriber listens for log messages and saves them to MongoDB.
type NatsSubscriber struct {
	natsClient    *NatsClient
	mongoClient   *mongo.MongoClient
	logCollection string
}

// NewNatsSubscriber creates a new instance of NatsSubscriber.
func NewNatsSubscriber(natsClient *NatsClient, mongoClient *mongo.MongoClient, logCollection string) *NatsSubscriber {
	return &NatsSubscriber{
		natsClient:    natsClient,
		mongoClient:   mongoClient,
		logCollection: logCollection,
	}
}

// SubscribeToLogs listens for log messages on the given NATS subject and saves them to MongoDB.
func (s *NatsSubscriber) SubscribeToLogs(subject string) error {
	return s.natsClient.Subscribe(subject, func(msg *nats.Msg) {
		var logEntry domain.LogEntry
		if err := json.Unmarshal(msg.Data, &logEntry); err != nil {
			log.Printf("Error unmarshalling log entry: %v", err)
			return
		}

		// Сохраняем лог в MongoDB
		if err := s.mongoClient.InsertLog(s.logCollection, logEntry); err != nil {
			log.Printf("Error saving log to MongoDB: %v", err)
			return
		}
		log.Printf("Log saved to MongoDB: %+v", logEntry)
	})
}
