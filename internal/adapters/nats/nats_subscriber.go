package nats

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"

	"github.com/bohexists/log-aggregator-svc/internal/domain"
	"github.com/bohexists/log-aggregator-svc/ports/inbound"
)

// NatsSubscriber listens for log messages and sends them to the consumer for further processing.
type NatsSubscriber struct {
	natsClient *NatsClient
	consumer   *inbound.NatsConsumer
}

// NewNatsSubscriber creates a new instance of NatsSubscriber.
func NewNatsSubscriber(natsClient *NatsClient, consumer *inbound.NatsConsumer) *NatsSubscriber {
	return &NatsSubscriber{
		natsClient: natsClient,
		consumer:   consumer,
	}
}

// SubscribeToLogs listens for log messages on the given NATS subject and sends them to the NatsConsumer for processing.
func (s *NatsSubscriber) SubscribeToLogs(subject string) error {
	// Используем существующий метод Subscribe из NatsClient
	return s.natsClient.Subscribe(subject, func(msg *nats.Msg) {
		var logEntry domain.LogEntry
		// Unmarshal the received message into a log entry
		if err := json.Unmarshal(msg.Data, &logEntry); err != nil {
			log.Printf("Error unmarshalling log entry: %v", err)
			return
		}

		// Send the log entry to NatsConsumer for processing
		if err := s.consumer.ConsumeLog(&logEntry); err != nil {
			log.Printf("Error processing log entry: %v", err)
			return
		}
		log.Printf("Log entry successfully processed: %+v", logEntry)
	})
}
