package logger

import (
	"os"
	"strings"

	"github.com/hueristiq/hq-go-logger/formatter"
	"github.com/hueristiq/hq-go-logger/levels"
	"github.com/hueristiq/hq-go-logger/writer"
)

// _Event represents a log event with a severity level, message, and optional metadata.
// It is used internally by the Logger to construct log messages before formatting and
// writing. The event is built using the options pattern, allowing flexible configuration
// of its fields via Option functions.
//
// Fields:
//   - level (levels.Level): The severity level of the log message, as defined in the
//     levels package (e.g., LevelInfo, LevelFatal).
//   - message (string): The primary content of the log message.
//   - metadata (map[string]string): Optional key-value pairs for additional context,
//     such as labels or system metrics. The "label" key is used for formatted output.
type _Event struct {
	level    levels.Level
	message  string
	metadata map[string]interface{}
}

// WithLevel sets the severity level of the log event.
//
// Parameters:
//   - level (levels.Level): The severity level to set, from the levels package.
func (e *_Event) WithLevel(level levels.Level) {
	e.level = level
}

// WithMessage sets the message content of the log event.
//
// Parameters:
//   - message (string): The log message to set.
func (e *_Event) WithMessage(message string) {
	e.message = message
}

// WithString adds a key-value pair to the log event's metadata.
//
// Parameters:
//   - key (string): The metadata key.
//   - value (string): The metadata value.
func (e *_Event) WithString(key, value string) {
	e.metadata[key] = value
}

// WithLabel sets the "label" metadata field for the log event, used by formatters
// to identify the log level in output (e.g., "[INFO]").
//
// Parameters:
//   - label (string): The label to set in the metadata.
func (e *_Event) WithLabel(label string) {
	e.WithString("label", label)
}

// WithError adds an error to the log event's metadata under the "error" key. The
// error value is stored as-is, and formatters are responsible for converting it to
// a string or other format as needed.
//
// Parameters:
//   - err (error): The error to set in the metadata.
func (e *_Event) WithError(err error) {
	e.metadata["error"] = err
}

// Logger is the main component of the logging system, responsible for filtering,
// formatting, and writing log messages. It uses a configured severity threshold to
// filter messages, a formatter to convert events to byte slices, and a writer to
// output the formatted data. The Logger provides level-specific methods (e.g., Info,
// Fatal) for convenient logging and supports metadata via the options pattern.
//
// Fields:
//   - level (levels.Level): The minimum severity level for logging (inclusive).
//     Messages with a higher level value (less severe) are ignored.
//   - formatter (formatter.Formatter): The formatter to convert log events to byte
//     slices.
//   - writer (writer.Writer): The writer to output formatted log data.
type Logger struct {
	level     levels.Level
	formatter formatter.Formatter
	writer    writer.Writer
}

// SetLevel sets the minimum severity level for logging. Messages with a level
// greater than the specified level (less severe) are ignored. The levels package
// uses lower values for higher severity (e.g., LevelFatal = 0, LevelDebug = 5).
//
// Parameters:
//   - level (levels.Level): The minimum severity level to log.
func (l *Logger) SetLevel(level levels.Level) {
	l.level = level
}

// SetFormatter sets the formatter used to convert log events to byte slices.
//
// Parameters:
//   - f (formatter.Formatter): The formatter to use for log events.
func (l *Logger) SetFormatter(f formatter.Formatter) {
	l.formatter = f
}

// SetWriter sets the writer used to output formatted log data.
//
// Parameters:
//   - w (writer.Writer): The writer to use for log output.
func (l *Logger) SetWriter(w writer.Writer) {
	l.writer = w
}

// Fatal logs a message at LevelFatal, applying the provided options (e.g., metadata).
// The message is formatted and written if the logger's threshold allows, and the
// program exits with status code 1 afterward. The levels package defines LevelFatal
// as the most severe level (value 0), so it is always logged unless the formatter or
// writer is nil.
//
// Parameters:
//   - message (string): The log message.
//   - options (...Option): Optional configurations for the log event (e.g., metadata).
func (l *Logger) Fatal(message string, options ...Option) {
	options = append(options, _WithLevel(levels.LevelFatal), _WithMessage(message))

	l.Log(_NewEvent(options...))
}

// Print logs a message at LevelSilent, applying the provided options. The message
// is formatted and written if the logger's threshold allows. LevelSilent (value 1)
// is typically used for non-critical output and may be directed to stdout by writers.
//
// Parameters:
//   - message (string): The log message.
//   - options (...Option): Optional configurations for the log event.
func (l *Logger) Print(message string, options ...Option) {
	options = append(options, _WithLevel(levels.LevelSilent), _WithMessage(message))

	l.Log(_NewEvent(options...))
}

// Error logs a message at LevelError, applying the provided options. The message
// is formatted and written if the logger's threshold allows (level <= LevelError).
// LevelError (value 2) indicates errors requiring attention but not program termination.
//
// Parameters:
//   - message (string): The log message.
//   - options (...Option): Optional configurations for the log event.
func (l *Logger) Error(message string, options ...Option) {
	options = append(options, _WithLevel(levels.LevelError), _WithMessage(message))

	l.Log(_NewEvent(options...))
}

// Info logs a message at LevelInfo, applying the provided options. The message
// is formatted and written if the logger's threshold allows (level <= LevelInfo).
// LevelInfo (value 3) is used for informational messages about normal operation.
//
// Parameters:
//   - message (string): The log message.
//   - options (...Option): Optional configurations for the log event.
func (l *Logger) Info(message string, options ...Option) {
	options = append(options, _WithLevel(levels.LevelInfo), _WithMessage(message))

	l.Log(_NewEvent(options...))
}

// Warn logs a message at LevelWarn, applying the provided options. The message
// is formatted and written if the logger's threshold allows (level <= LevelWarn).
// LevelWarn (value 4) indicates potential issues that do not halt execution.
//
// Parameters:
//   - message (string): The log message.
//   - options (...Option): Optional configurations for the log event.
func (l *Logger) Warn(message string, options ...Option) {
	options = append(options, _WithLevel(levels.LevelWarn), _WithMessage(message))

	l.Log(_NewEvent(options...))
}

// Debug logs a message at LevelDebug, applying the provided options. The message
// is formatted and written if the logger's threshold allows (level <= LevelDebug).
// LevelDebug (value 5) is used for detailed debugging information.
//
// Parameters:
//   - message (string): The log message.
//   - options (...Option): Optional configurations for the log event.
func (l *Logger) Debug(message string, options ...Option) {
	options = append(options, _WithLevel(levels.LevelDebug), _WithMessage(message))

	l.Log(_NewEvent(options...))
}

// Log processes a log event by filtering, formatting, and writing it. The event is
// ignored if its level is greater than the logger's threshold (less severe). If no
// label is provided in the event's metadata, a default label is added (e.g., "INF"
// for LevelInfo). The message is trimmed of trailing newlines before formatting.
// If the formatter or writer is nil, or if formatting fails, the event is silently
// ignored. For LevelFatal events, the program exits with status code 1 after writing.
//
// Parameters:
//   - event (*_Event): The log event to process, containing level, message, and metadata.
func (l *Logger) Log(event *_Event) {
	if event.level > l.level {
		return
	}

	var (
		ok    bool
		label string
	)

	if _, ok = event.metadata["label"]; !ok {
		labels := map[levels.Level]string{
			levels.LevelFatal: "FTL",
			levels.LevelError: "ERR",
			levels.LevelInfo:  "INF",
			levels.LevelWarn:  "WRN",
			levels.LevelDebug: "DBG",
		}

		if label, ok = labels[event.level]; ok {
			event.metadata["label"] = label
		}
	}

	event.message = strings.TrimSuffix(event.message, "\n")

	data, err := l.formatter.Format(&formatter.Log{
		Message:  event.message,
		Level:    event.level,
		Metadata: event.metadata,
	})
	if err != nil {
		return
	}

	l.writer.Write(data, event.level)

	if event.level == levels.LevelFatal {
		os.Exit(1)
	}
}

// Option defines a function type for configuring log events using the options pattern.
//
// Parameters:
//   - event (*_Event): The log event to configure
type Option func(*_Event)

// _NewEvent creates a new log event with the specified options. It initializes the
// event with an empty metadata map and applies the provided options to set the level,
// message, and metadata.
//
// Parameters:
//   - options (...Option): Configurations for the log event (e.g., level, message, metadata).
//
// Returns:
//   - event (*_Event): A pointer to the configured log event.
func _NewEvent(options ...Option) (event *_Event) {
	event = &_Event{
		metadata: make(map[string]interface{}),
	}

	for _, option := range options {
		option(event)
	}

	return
}

// _WithLevel returns an Option that sets the severity level of a log event.
//
// Parameters:
//   - level (levels.Level): The severity level to set.
//
// Returns:
//   - option (Option): A function to configure the event's level.
func _WithLevel(level levels.Level) (option Option) {
	return func(event *_Event) {
		event.WithLevel(level)
	}
}

// _WithMessage returns an Option that sets the message content of a log event.
//
// Parameters:
//   - message (string): The log message to set.
//
// Returns:
//   - option (Option): A function to configure the event's message.
func _WithMessage(message string) (option Option) {
	return func(event *_Event) {
		event.WithMessage(message)
	}
}

// WithString returns an Option that adds a key-value pair to a log event's metadata.
//
// Parameters:
//   - key (string): The metadata key.
//   - value (string): The metadata value.
//
// Returns:
//   - option (Option): A function to configure the event's metadata.
func WithString(key, value string) (option Option) {
	return func(event *_Event) {
		event.WithString(key, value)
	}
}

// WithLabel returns an Option that sets the "label" metadata field for a log event.
//
// Parameters:
//   - label (string): The label to set in the metadata.
//
// Returns:
//   - option (Option): A function to configure the event's label.
func WithLabel(label string) (option Option) {
	return func(event *_Event) {
		event.WithLabel(label)
	}
}

// WithError returns an Option that adds an error to a log event's metadata under
// the "error" key. The error is stored as-is, and formatters are responsible for
// converting it to a string or other format as needed.
//
// Parameters:
//   - err (error): The error to set in the metadata.
//
// Returns:
//   - option (Option): A function to configure the event's error metadata.
func WithError(err error) (option Option) {
	return func(event *_Event) {
		event.WithError(err)
	}
}

// NewLogger creates and returns a new Logger instance with default settings (no
// formatter, writer, or level set). Users must configure the logger with a level,
// formatter, and writer before use to avoid silent failures during logging.
//
// Returns:
//   - logger (*Logger): A pointer to a new Logger instance.
func NewLogger() (logger *Logger) {
	logger = &Logger{}

	return
}
