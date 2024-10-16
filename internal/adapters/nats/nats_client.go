package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

type NatsClient struct {
	conn *nats.Conn
}

// NewNatsClient creates a new NatsClient instance.
func NewNatsClient(natsURL string) (*NatsClient, error) {
	conn, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	return &NatsClient{conn: conn}, nil
}

// Publish publishes a message to NATS.
func (n *NatsClient) Publish(subject string, data []byte) error {
	return n.conn.Publish(subject, data)
}

// Subscribe subscribes to a NATS subject and handles messages.
func (n *NatsClient) Subscribe(subject string, handler nats.MsgHandler) error {
	_, err := n.conn.Subscribe(subject, handler)
	if err != nil {
		return err
	}
	log.Printf("Subscribed to subject: %s", subject)
	return nil
}
