package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/hueristiq/hq-go-logger/formatter"
	"github.com/hueristiq/hq-go-logger/levels"
	"github.com/hueristiq/hq-go-logger/writer"
)

type Logger struct {
	formatter formatter.Formatter
	level     levels.Level
	writer    writer.Writer
}

// SetFormatter updates the Formatter used by the Logger.
//
// Parameters:
//
//   - f (formatter.Formatter): The formatting strategy (e.g., console, JSON).
func (l *Logger) SetFormatter(f formatter.Formatter) {
	l.formatter = f
}

// SetLevel sets the logging threshold. Events with a level greater than this are no-ops.
//
// Parameters:
//
//   - level (levels.Level): Maximum level to emit (lower numeric means higher severity).
func (l *Logger) SetLevel(level levels.Level) {
	l.level = level
}

// SetWriter specifies where formatted log entries are written.
//
// Parameters:
//
//   - w (writer.Writer): The output target (e.g., console writer, file writer).
func (l *Logger) SetWriter(w writer.Writer) {
	l.writer = w
}

// Log processes a completed Event: it filters by level, applies default labels,
// trims trailing newlines, formats via Formatter, writes via Writer, and on Fatal,
// exits the program.
//
// Steps:
// 1. Suppress if event.level > Logger.level.
// 2. If no "label" metadata, assign default codes (FTL, ERR, INF, WRN, DBG).
// 3. Trim any trailing '\n' from message.
// 4. Format the event: produce []byte or abort on error.
// 5. Write the data and level via Writer.
// 6. If level == LevelFatal, call os.Exit(1).
//
// Parameters:
//
//   - event (*Event): The fully-built event to log.
func (l *Logger) Log(event *Event) {
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

func (l *Logger) Fatal() (event *Event) {
	event = &Event{
		logger:   l,
		level:    levels.LevelFatal,
		metadata: make(map[string]string),
	}

	return
}

func (l *Logger) Print() (event *Event) {
	event = &Event{
		logger:   l,
		level:    levels.LevelSilent,
		metadata: make(map[string]string),
	}

	return
}

func (l *Logger) Error() (event *Event) {
	event = &Event{
		logger:   l,
		level:    levels.LevelError,
		metadata: make(map[string]string),
	}

	return
}

func (l *Logger) Info() (event *Event) {
	event = &Event{
		logger:   l,
		level:    levels.LevelInfo,
		metadata: make(map[string]string),
	}

	return
}

func (l *Logger) Warn() (event *Event) {
	event = &Event{
		logger:   l,
		level:    levels.LevelWarn,
		metadata: make(map[string]string),
	}

	return
}

func (l *Logger) Debug() (event *Event) {
	event = &Event{
		logger:   l,
		level:    levels.LevelDebug,
		metadata: make(map[string]string),
	}

	return
}

// Event is a builder for a log entry; it captures level, message, and metadata
// before passing itself to the Logger for actual writing.
//
// Typical usage: logger.Info().Label("HTTP").Msg("Request received").
//
// Fields:
//
//   - logger   (*Logger)        - Back-reference to the parent Logger for invoking Log().
//   - level    (levels.Level)   - Severity level of the event.
//   - message  (string)         - Log message content.
//   - metadata (map[string]string) - Arbitrary key-value pairs for context.
type Event struct {
	logger *Logger

	level    levels.Level
	message  string
	metadata map[string]string
}

// Label sets the "label" metadata for the Event, e.g. component or category tag.
//
// Parameters:
//
//   - label (string): Short identifier to include in the formatted output.
//
// Returns:
//
//	-event *Event: Same Event to allow chaining.
func (e *Event) Label(label string) (event *Event) {
	e.metadata["label"] = label

	return e
}

// Msg finalizes the Event with a literal message and dispatches it to the Logger.
//
// Parameters:
//
//   - message (string): The content to log (may include format verbs if using Msgf).
func (e *Event) Msg(message string) {
	e.message = message

	e.logger.Log(e)
}

// Msgf formats the Event message using fmt.Sprintf and dispatches it.
//
// Parameters:
//
//   - format (string): A fmt.Sprintf format string.
//   - args   (...interface{}): Arguments for the format verbs.
func (e *Event) Msgf(format string, args ...interface{}) {
	e.message = fmt.Sprintf(format, args...)

	e.logger.Log(e)
}

var DefaultLogger *Logger

func init() {
	DefaultLogger = &Logger{}

	DefaultLogger.SetFormatter(formatter.NewConsoleFormatter(&formatter.ConsoleFormatterConfiguration{
		Colorize: true,
	}))
	DefaultLogger.SetLevel(levels.LevelDebug)
	DefaultLogger.SetWriter(writer.NewConsoleWriter())
}

func Fatal() (event *Event) {
	event = &Event{
		logger:   DefaultLogger,
		level:    levels.LevelFatal,
		metadata: make(map[string]string),
	}

	return
}

func Print() (event *Event) {
	event = &Event{
		logger:   DefaultLogger,
		level:    levels.LevelSilent,
		metadata: make(map[string]string),
	}

	return
}

func Error() (event *Event) {
	event = &Event{
		logger:   DefaultLogger,
		level:    levels.LevelError,
		metadata: make(map[string]string),
	}

	return
}

func Info() (event *Event) {
	event = &Event{
		logger:   DefaultLogger,
		level:    levels.LevelInfo,
		metadata: make(map[string]string),
	}

	return
}

func Warn() (event *Event) {
	event = &Event{
		logger:   DefaultLogger,
		level:    levels.LevelWarn,
		metadata: make(map[string]string),
	}

	return
}

func Debug() (event *Event) {
	event = &Event{
		logger:   DefaultLogger,
		level:    levels.LevelDebug,
		metadata: make(map[string]string),
	}

	return
}
