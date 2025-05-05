package writer

import (
	"os"
	"sync"

	"github.com/hueristiq/hq-go-logger/levels"
)

// Console is a Writer implementation that sends log entries to the console (stdout or stderr).
//
// It uses an internal mutex to serialize writes from multiple goroutines, preventing
// interleaved output. Messages at LevelSilent are written to stdout; all other levels
// are written to stderr. Each call to Write appends a newline after the data.
//
// Fields:
//
// - mutex (*sync.Mutex) - Protects writes to the console to ensure atomicity.
type Console struct {
	mutex *sync.Mutex
}

// Write sends the provided log data to an appropriate output stream based on the log level.
// It locks the writer's mutex to ensure that writes from concurrent goroutines do not overlap.
//
// The output routing is determined by the log level:
//   - For LevelSilent, the data is written to os.Stdout.
//   - For all other levels, the data is written to os.Stderr.
//
// After writing the data, a newline character is appended to the output stream.
//
// Parameters:
//
// - data ([]byte): A byte slice containing the formatted log message.
// - level (levels.Level): The severity level of the log message, used to determine the output destination.
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

// NewConsoleWriter creates and returns a new Console writer instance.
// It initializes the internal mutex to guarantee safe concurrent writes.
//
// Returns:
//
// - writer (*Console): A pointer to a Console writer, ready for use in writing log messages.
func NewConsoleWriter() (writer *Console) {
	writer = &Console{
		mutex: &sync.Mutex{},
	}

	return
}
