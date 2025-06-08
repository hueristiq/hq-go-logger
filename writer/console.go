package writer

import (
	"os"
	"sync"

	"github.com/hueristiq/hq-go-logger/levels"
)

// Console is a thread-safe implementation of the Writer interface that writes log
// messages to standard output (stdout) or standard error (stderr) based on the log
// level.
//
// Fields:
//   - mutex (*sync.Mutex): Ensures thread-safe access to stdout and stderr during
//     write operations.
type Console struct {
	mutex *sync.Mutex
}

// Write writes the provided log data to either stdout or stderr based on the
// specified log level, appending a newline character to each message. Messages with
// LevelSilent are written to stdout, while all other levels (LevelFatal, LevelError,
// LevelInfo, LevelWarn, LevelDebug) are written to stderr. The method is thread-safe,
// using a mutex to prevent concurrent write conflicts to the output streams.
//
// Parameters:
//   - data ([]byte): The pre-formatted log message to write, typically produced by
//     a formatter (e.g., as JSON or plain text).
//   - level (levels.Level): The severity level of the log message, as defined in
//     the levels package.
func (c *Console) Write(data []byte, level levels.Level) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	switch level {
	case levels.LevelSilent:
		os.Stdout.Write(data)
		os.Stdout.WriteString("\n")
	default:
		os.Stderr.Write(data)
		os.Stderr.WriteString("\n")
	}
}

var _ Writer = (*Console)(nil)

// NewConsoleWriter creates and returns a new Console writer instance, initialized
// with a mutex for thread-safe operation.
//
// Returns:
//   - writer (*Console): A pointer to a new Console writer instance, ready for use
//     in a logging system.
func NewConsoleWriter() (writer *Console) {
	writer = &Console{
		mutex: &sync.Mutex{},
	}

	return
}
