package nats

import (
	"encoding/json"
	"log"

	"github.com/bohexists/log-aggregator-svc/internal/domain"
	"github.com/bohexists/log-aggregator-svc/internal/services"
	"github.com/nats-io/nats.go"
)

// NatsSubscriber listens for log messages and sends them to the log service for further processing.
type NatsSubscriber struct {
	natsClient *NatsClient
	logService *services.LogService
}

// NewNatsSubscriber creates a new instance of NatsSubscriber.
func NewNatsSubscriber(natsClient *NatsClient, logService *services.LogService) *NatsSubscriber {
	return &NatsSubscriber{
		natsClient: natsClient,
		logService: logService,
	}
}

// SubscribeToLogs listens for log messages on the given NATS subject and sends them to the LogService for processing.
func (s *NatsSubscriber) SubscribeToLogs(subject string) error {
	return s.natsClient.Subscribe(subject, func(msg *nats.Msg) {
		var logEntry domain.LogEntry
		// Разбираем сообщение из NATS
		if err := json.Unmarshal(msg.Data, &logEntry); err != nil {
			log.Printf("Error unmarshalling log entry: %v", err)
			return
		}

		// Передаем разобранные данные в бизнес-логику
		if err := s.logService.ProcessLog(logEntry); err != nil {
			log.Printf("Error processing log entry: %v", err)
			return
		}
		log.Printf("Log entry successfully processed: %+v", logEntry)
	})
}
