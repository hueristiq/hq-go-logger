package formatter

import "github.com/hueristiq/hq-go-logger/levels"

// Log represents a structured log message.
// It encapsulates the core message, its severity level, and any additional metadata that provides context.
//
// Fields:
//
//	Message  - The primary content of the log message.
//	Level    - The severity level of the log message (e.g., Debug, Info, Warn, Error, Fatal),
//	           represented by a value from the levels package.
//	Metadata - A map containing additional contextual key-value pairs, such as labels or identifiers.
type Log struct {
	Message  string
	Level    levels.Level
	Metadata map[string]string
}

// Formatter defines the interface that must be implemented by any log formatter.
// A Formatter converts a Log instance into a formatted output in the form of a byte slice.
// This abstraction allows for multiple formatting strategies (e.g., console, file, JSON) to coexist.
type Formatter interface {
	Format(log *Log) (data []byte, err error)
}
