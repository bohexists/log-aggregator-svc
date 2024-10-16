package config

import (
	"os"
)

// Config struct holds all the necessary configuration variables.
type Config struct {
	NatsURL  string
	MongoURL string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() Config {
	return Config{
		NatsURL:  os.Getenv("NATS_URL"),
		MongoURL: os.Getenv("MONGO_URL"),
	}
}
