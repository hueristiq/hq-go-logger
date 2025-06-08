package writer

import "github.com/hueristiq/hq-go-logger/levels"

// Writer defines the interface for writing log messages to an output destination.
// Implementations of this interface handle the delivery of formatted log data to
// specific sinks, such as files, consoles, network endpoints, or external logging
// services, while considering the severity level of the log message.
//
// Methods:
//   - Write(data []byte, level levels.Level): Writes the provided log data to the
//     output destination if the severity level meets the writer's criteria (e.g., a
//     configured threshold). The data is typically pre-formatted by a formatter
//     (e.g., as JSON or plain text), and the level is a severity value from the
//     levels package (e.g., LevelInfo, LevelError). No return value is required,
//     as errors are typically handled internally (e.g., by logging to a fallback
//     or discarding the message).
type Writer interface {
	Write(data []byte, level levels.Level)
}
