package logger

import (
	"github.com/hueristiq/hq-go-logger/formatter"
	"github.com/hueristiq/hq-go-logger/levels"
	"github.com/hueristiq/hq-go-logger/writer"
)

// DefaultLogger is a pre-configured Logger instance for convenient logging without
// explicit instantiation. It is initialized in the init() function with:
//   - Level: LevelDebug (allows all messages).
//   - Formatter: Console formatter with colorized labels (Colorize: true).
//   - Writer: Console writer directing LevelSilent to stdout and other levels to stderr.
//
// Users can log directly using package-level functions (Fatal, Print, Error, Info,
// Warn, Debug), which delegate to DefaultLogger. The Logger filters messages based
// on its level threshold, adds default labels if none are provided, and exits the
// program with status code 1 for LevelFatal messages. The levels package uses lower
// values for higher severity (e.g., LevelFatal = 0, LevelDebug = 5).
//
// To customize logging behavior (e.g., change level, formatter, or writer), users
// can modify DefaultLogger's configuration or create a new Logger instance.
var DefaultLogger *Logger

func init() {
	DefaultLogger = NewLogger()

	DefaultLogger.SetLevel(levels.LevelDebug)
	DefaultLogger.SetFormatter(formatter.NewConsoleFormatter(&formatter.ConsoleFormatterConfiguration{
		Colorize: true,
	}))
	DefaultLogger.SetWriter(writer.NewConsoleWriter())
}

// Fatal logs a message at LevelFatal using DefaultLogger, applying the provided
// options (e.g., metadata). The message is formatted and written if the logger's
// threshold allows, and the program exits with status code 1 afterward. LevelFatal
// (value 0) is the most severe level, so it is always logged unless the formatter
// or writer is nil.
//
// Parameters:
//   - message (string): The log message.
//   - ofs (...OptionFunc): Optional configurations for the log event (e.g., metadata).
func Fatal(message string, ofs ...OptionFunc) {
	DefaultLogger.Fatal(message, ofs...)
}

// Print logs a message at LevelSilent using DefaultLogger, applying the provided
// options. The message is formatted and written if the logger's threshold allows.
// LevelSilent (value 1) is typically used for non-critical output and may be
// directed to stdout by writers (e.g., the default Console writer).
//
// Parameters:
//   - message (string): The log message.
//   - ofs (...OptionFunc): Optional configurations for the log event.
func Print(message string, ofs ...OptionFunc) {
	DefaultLogger.Print(message, ofs...)
}

// Error logs a message at LevelError using DefaultLogger, applying the provided
// options. The message is formatted and written if the logger's threshold allows
// (level <= LevelError). LevelError (value 2) indicates errors requiring attention
// but not program termination.
//
// Parameters:
//   - message (string): The log message.
//   - ofs (...OptionFunc): Optional configurations for the log event.
func Error(message string, ofs ...OptionFunc) {
	DefaultLogger.Error(message, ofs...)
}

// Info logs a message at LevelInfo using DefaultLogger, applying the provided
// options. The message is formatted and written if the logger's threshold allows
// (level <= LevelInfo). LevelInfo (value 3) is used for informational messages
// about normal operation.
//
// Parameters:
//   - message (string): The log message.
//   - ofs (...OptionFunc): Optional configurations for the log event.
func Info(message string, ofs ...OptionFunc) {
	DefaultLogger.Info(message, ofs...)
}

// Warn logs a message at LevelWarn using DefaultLogger, applying the provided
// options. The message is formatted and written if the logger's threshold allows
// (level <= LevelWarn). LevelWarn (value 4) indicates potential issues that do not
// halt execution.
//
// Parameters:
//   - message (string): The log message.
//   - ofs (...OptionFunc): Optional configurations for the log event.
func Warn(message string, ofs ...OptionFunc) {
	DefaultLogger.Warn(message, ofs...)
}

// Debug logs a message at LevelDebug using DefaultLogger, applying the provided
// options. The message is formatted and written if the logger's threshold allows
// (level <= LevelDebug). LevelDebug (value 5) is used for detailed debugging
// information.
//
// Parameters:
//   - message (string): The log message.
//   - ofs (...OptionFunc): Optional configurations for the log event.
func Debug(message string, ofs ...OptionFunc) {
	DefaultLogger.Debug(message, ofs...)
}
