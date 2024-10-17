package outbound

import "github.com/bohexists/log-aggregator-svc/internal/domain"

// LogRepository defines the contract for log data storage.
type LogRepository interface {
	// InsertLog saves a log entry to the repository.
	InsertLog(logEntry *domain.LogEntry) error

	// Close gracefully disconnects from the repository.
	Close() error
}
