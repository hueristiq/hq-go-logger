package formatter

import (
	"bytes"

	"github.com/logrusorgru/aurora/v4"
	"go.source.hueristiq.com/logger/levels"
)

// Console implements the Formatter interface for formatting logs intended for console output.
// It uses an aurora.Aurora instance to apply colorization effects (if enabled) to log metadata labels.
type Console struct {
	au *aurora.Aurora
}

// Format constructs the final log message as a byte slice ready for output on the console.
// It first applies colorization to the log's label (if available), then builds the formatted message.
// The final message consists of an optional label (enclosed in square brackets and followed by a space)
// and the main log message.
//
// Parameters:
//
//	log (*Log): a pointer to a Log structure containing the log message and its metadata.
//
// Returns:
//
//	data ([]byte): a byte slice containing the formatted log message.
//	err (error): an error value, which is always nil in the current implementation.
func (c *Console) Format(log *Log) (data []byte, err error) {
	c.colorize(log)

	buffer := &bytes.Buffer{}

	buffer.Grow(len(log.Message))

	if label, ok := log.Metadata["label"]; ok && label != "" {
		buffer.WriteByte('[')
		buffer.WriteString(label)
		buffer.WriteByte(']')
		buffer.WriteByte(' ')
	}

	buffer.WriteString(log.Message)

	data = buffer.Bytes()

	return
}

// colorize applies ANSI color codes to the log's "label" metadata based on the log level.
// If a label is present, it is transformed into a colorized and bold version.
// The color is selected according to the log level as follows:
//   - Fatal and Error levels: Bright Red
//   - Info level: Bright Blue
//   - Warn level: Bright Yellow
//   - Debug level: Bright Magenta
//
// If no label is provided, the function exits without modifying the log.
// Note: The aurora library is used to wrap the label with the proper ANSI escape sequences.
func (c *Console) colorize(log *Log) {
	label := log.Metadata["label"]

	if label == "" {
		return
	}

	//nolint: exhaustive
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

// ConsoleFormatterConfiguration holds configuration options for the Console formatter.
// It allows the caller to enable or disable ANSI colorization of the log label.
type ConsoleFormatterConfiguration struct {
	Colorize bool
}

var _ Formatter = (*Console)(nil)

// NewConsoleFormatter creates and returns a new instance of the Console formatter.
// It initializes the aurora.Aurora instance using the provided configuration to determine whether
// to enable colorization.
//
// Parameters:
//
//	cfg (*ConsoleFormatterConfiguration): a pointer to a ConsoleFormatterConfiguration specifying the colorization option.
//
// Returns:
//
//	formatter (*Console): A pointer to a new Console formatter instance.
func NewConsoleFormatter(cfg *ConsoleFormatterConfiguration) (formatter *Console) {
	formatter = &Console{
		au: aurora.New(aurora.WithColors(cfg.Colorize)),
	}

	return
}
