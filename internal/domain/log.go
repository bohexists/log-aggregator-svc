package domain

import (
	"time"
)

// LogEntry represents a log entry in the system.
type LogEntry struct {
	ID        string    `bson:"_id,omitempty"`
	Message   string    `bson:"message"`
	Timestamp time.Time `bson:"timestamp"`
	Type      string    `bson:"type"`
}

// NewLogEntry creates a new log entry.
func NewLogEntry(message, logType string) *LogEntry {
	return &LogEntry{
		Message:   message,
		Timestamp: time.Now(),
		Type:      logType,
	}
}
