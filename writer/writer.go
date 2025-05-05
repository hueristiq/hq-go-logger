package writer

import "github.com/hueristiq/hq-go-logger/levels"

// Writer is the interface for writing serialized log output to a target system.
// Implementations decide how and where to send the byte slice representing the log entry,
// and may use the Level parameter for conditional behavior (e.g., separate files per level,
// dynamic log routing, or filtering.)
//
// Implementers must ensure that Write operations are safe for concurrent use if the Writer
// will be shared across goroutines. Consider buffering, batching, or synchronization as needed.
//
// Example:
//
//	type FileWriter struct {
//	    mu   sync.Mutex
//	    file *os.File
//	}
//
//	func (w *FileWriter) Write(data []byte, level levels.Level) {
//	    w.mu.Lock()
//	    defer w.mu.Unlock()
//	    // Example: prefix each line with level name
//	    line := fmt.Sprintf("%s %s", level.String(), string(data))
//	    w.file.Write([]byte(line))
//	}
//
// Example usage with a logging system:
//
//	var writer writer.Writer = &FileWriter{file: os.Stdout}
//
//	data, _ := consoleFormatter.Format(logEntry)
//
//	writer.Write(data, logEntry.Level)
type Writer interface {
	// Write sends the formatted log data to the target output.
	//
	// Parameters:
	//
	//   - data ([]byte): Serialized log entry (e.g., text, JSON) produced by a Formatter.
	//   - level (levels.Level): Severity level of the log entry. Can be used for routing or filtering.
	Write(data []byte, level levels.Level)
}
