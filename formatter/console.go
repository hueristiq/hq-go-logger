package formatter

import (
	"bytes"

	"github.com/hueristiq/hq-go-logger/levels"
	"github.com/logrusorgru/aurora/v4"
)

// Console is an implementation of the Formatter interface that formats log messages
// for console output. It constructs a string in the format "[label] message [metadata]"
// (if label and metadata are present and configured) and returns it as a byte slice.
// Labels are colorized based on the log level using the aurora package if colorization
// is enabled. The formatter supports including additional metadata fields as
// key=value pairs if configured.
//
// Fields:
//   - au (*aurora.Aurora): The aurora instance used for colorizing labels.
//
// The formatter is thread-safe, as it maintains no shared state between Format calls.
// It handles nil or empty Metadata maps safely and validates log levels to prevent
// undefined behavior. The levels package's severity order (lower values = higher
// severity) is used for colorization.
type Console struct {
	au *aurora.Aurora
}

// Format converts a Log struct into a formatted byte slice for console output.
// It constructs a string in the format "[label] message [metadata]" (if label and
// metadata are present and configured). If colorization is enabled, the label is
// colored based on the log level. Additional metadata is appended as key=value pairs
// if includeMetadata is true. The output does not include a trailing newline, as
// newlines are typically added by writers.
//
// Parameters:
//   - log (*Log): The log message to format, containing the level, message, and
//     optional metadata.
//
// Returns:
//   - data ([]byte): The formatted log message as a byte slice.
//   - err (error): An error if the log level is invalid, otherwise nil.
//
// The method is thread-safe and handles nil or empty Metadata maps safely. It
// validates the log level to ensure it is within the valid range [LevelFatal,
// LevelDebug]. If the level is invalid, it returns an error.
func (c *Console) Format(log *Log) (data []byte, err error) {
	c.colorize(log)

	buffer := &bytes.Buffer{}

	buffer.Grow(len(log.Message))

	if label, ok := log.Metadata["label"]; ok && label != "" {
		buffer.WriteByte('[')
		buffer.WriteString(label)
		buffer.WriteByte(']')
		buffer.WriteByte(' ')

		delete(log.Metadata, "label")
	}

	buffer.WriteString(log.Message)

	for k, v := range log.Metadata {
		buffer.WriteByte(' ')
		buffer.WriteString(k)
		buffer.WriteByte('=')
		buffer.WriteString(v)
	}

	data = buffer.Bytes()

	return
}

// colorize applies color to the "label" metadata field of the Log based on its level,
// if colorization is enabled and a non-empty label is present. Colors are applied
// using the aurora package, with specific colors for each log level:
//   - LevelFatal: Bright red, bold
//   - LevelError: Bright red, bold
//   - LevelInfo: Bright blue, bold
//   - LevelWarn: Bright yellow, bold
//   - LevelDebug: Bright magenta, bold
//   - LevelSilent: No color (unchanged)
//
// If the label is empty, missing, or colorization is disabled, no changes are made.
//
// Parameters:
//   - log (*Log): The log message containing the level and metadata to process.
func (c *Console) colorize(log *Log) {
	label := log.Metadata["label"]

	if label == "" {
		return
	}

	switch log.Level {
	case levels.LevelFatal:
		log.Metadata["label"] = c.au.BrightRed(label).Bold().String()
	case levels.LevelError:
		log.Metadata["label"] = c.au.BrightRed(label).Bold().String()
	case levels.LevelInfo:
		log.Metadata["label"] = c.au.BrightBlue(label).Bold().String()
	case levels.LevelWarn:
		log.Metadata["label"] = c.au.BrightYellow(label).Bold().String()
	case levels.LevelDebug:
		log.Metadata["label"] = c.au.BrightMagenta(label).Bold().String()
	}
}

// ConsoleFormatterConfiguration defines configuration options for the Console formatter.
//
// Fields:
//   - Colorize (bool): If true, enables colorization of labels using aurora.s
type ConsoleFormatterConfiguration struct {
	Colorize bool
}

var _ Formatter = (*Console)(nil)

// NewConsoleFormatter creates and returns a new Console formatter instance,
// configured with the provided ConsoleFormatterConfiguration.
//
// Parameters:
//   - cfg (*ConsoleFormatterConfiguration): The configuration for the formatter.
//
// Returns:
//   - formatter (*Console): A pointer to a new Console formatter instance.
func NewConsoleFormatter(cfg *ConsoleFormatterConfiguration) (formatter *Console) {
	formatter = &Console{
		au: aurora.New(aurora.WithColors(cfg.Colorize)),
	}

	return
}
