package formatter

import "github.com/hueristiq/hq-go-logger/levels"

// Log represents a structured log message with severity, content, and contextual metadata.
//
// Fields:
//
//   - Level (levels.Level): The severity level of the log entry, indicating its importance (e.g., Debug, Info, Warn, Error, Fatal).
//   - Message (string): The primary textual content of the log entry.
//   - Metadata (map[string]string): A map of arbitrary key-value pairs providing additional context,
//     such as request IDs, user identifiers, or component names.
//
// Example usage:
//
//	logEntry := &formatter.Log{
//	    Level:    levels.LevelInfo,
//	    Message:  "User login successful",
//	    Metadata: map[string]string{"user_id": "12345", "ip": "192.168.1.100"},
//	}
//
// When passed to a Formatter, this Log will be converted into the desired output format.
type Log struct {
	Level    levels.Level
	Message  string
	Metadata map[string]string
}

// Formatter defines the interface for converting a Log into a formatted byte slice.
// Implementations of this interface can serialize logs as human-readable text,
// machine-readable JSON, or any other custom format.
//
// The Format method should:
//   - Respect the log Level (e.g., include the level name in the output).
//   - Include the Message as the core content.
//   - Serialize Metadata as part of the output (order is implementation-defined).
//   - Return an error if formatting fails (for example, JSON marshaling errors).
//
// Example:
//
//	type JSONFormatter struct {}
//
//	func (f *JSONFormatter) Format(log *formatter.Log) ([]byte, error) {
//	    payload := map[string]interface{}{
//	        "level":    log.Level.String(),
//	        "message":  log.Message,
//	        "metadata": log.Metadata,
//	    }
//	    return json.Marshal(payload)
//	}
//
// Example usage with a logger:
//
//	var formatter formatter.Formatter = &JSONFormatter{}
//
//	data, err := formatter.Format(logEntry)
//	if err == nil {
//	    os.Stdout.Write(data)
//	}
//
// Formatter implementations should strive for thread safety if used concurrently by multiple goroutines.
type Formatter interface {
	// Format converts the given Log into a slice of bytes representing the formatted output.
	//
	// Parameters:
	//
	//   - log (*Log): Pointer to the Log instance to format.
	//
	// Returns:
	//
	//   - data ([]byte): The formatted representation of the log entry.
	//   - err (error) : Non-nil if an error occurred during formatting.
	Format(log *Log) (data []byte, err error)
}
