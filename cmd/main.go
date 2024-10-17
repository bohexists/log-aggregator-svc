package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bohexists/log-aggregator-svc/internal/adapters/mongo"
	"github.com/bohexists/log-aggregator-svc/internal/adapters/nats"
	"github.com/bohexists/log-aggregator-svc/internal/config"
	"github.com/bohexists/log-aggregator-svc/ports/inbound"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize NATS client
	natsClient, err := nats.NewNatsClient(cfg.NatsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	// Initialize MongoDB client
	mongoClient, err := mongo.NewMongoClient(cfg.MongoURL, cfg.MongoDB)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize NatsConsumer with Mongo repository
	consumer := inbound.NewNatsConsumer(mongoClient, "logs")

	// Initialize NATS subscriber with the consumer
	natsSubscriber := nats.NewNatsSubscriber(natsClient, consumer)

	// Subscribe to NATS subject "logs"
	err = natsSubscriber.SubscribeToLogs("logs")
	if err != nil {
		log.Fatalf("Failed to subscribe to NATS subject: %v", err)
	}

	// Graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("Shutting down...")
}
