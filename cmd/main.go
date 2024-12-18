package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bohexists/log-aggregator-svc/internal/adapters/mongo"
	"github.com/bohexists/log-aggregator-svc/internal/adapters/nats"
	"github.com/bohexists/log-aggregator-svc/internal/config"
	"github.com/bohexists/log-aggregator-svc/internal/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize NATS client
	natsClient, err := nats.NewNatsClient(cfg.NatsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	// Initialize MongoDB client and repository
	mongoClient, err := mongo.NewMongoClient(cfg.MongoURL, cfg.MongoDB)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	mongoRepo := mongo.NewMongoRepository(mongoClient)

	// Initialize log service
	logService := services.NewLogService(mongoRepo)

	// Initialize NATS subscriber with log service as handler
	natsSubscriber := nats.NewNatsSubscriber(natsClient, logService)

	// Subscribe to NATS subject "logs" and process incoming logs
	err = natsSubscriber.SubscribeToLogs("logs")
	if err != nil {
		log.Fatalf("Failed to subscribe and process logs: %v", err)
	}

	// Graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("Shutting down...")
}
