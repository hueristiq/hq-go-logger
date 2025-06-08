package formatter

import "github.com/hueristiq/hq-go-logger/levels"

// Log represents a single log message with its associated severity level, content,
// and optional metadata. It is used as the input to Formatter implementations to
// produce formatted output for logging systems.
//
// Fields:
//   - Level (levels.Level): The severity of the log message, as defined in the
//     levels package (e.g., LevelFatal, LevelInfo, LevelDebug). Lower values
//     indicate higher severity.
//   - Message (string): The primary content of the log message, describing the
//     event or condition being logged.
//   - Metadata (map[string]string): Optional key-value pairs providing additional
//     context for the log message, such as request IDs, user IDs, or system metrics.
type Log struct {
	Level    levels.Level
	Message  string
	Metadata map[string]interface{}
}

// Formatter defines the interface for formatting log messages. Implementations
// of this interface convert a Log struct into a byte slice, enabling output in
// various formats (e.g., JSON, plain text, or structured logging formats).
// This allows logging systems to support multiple output styles while maintaining
// a consistent input structure.
//
// Methods:
//   - Format(log *Log) (data []byte, err error): Converts the provided Log into
//     a byte slice representing the formatted log message. Returns an error if
//     formatting fails (e.g., due to serialization issues). The returned data
//     should be suitable for writing to an output destination, such as a file
//     or network stream.
type Formatter interface {
	Format(log *Log) (data []byte, err error)
}
