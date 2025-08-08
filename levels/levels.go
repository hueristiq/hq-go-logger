package levels

import (
	"errors"
	"fmt"
)

// Level represents the severity of a log message. It is an integer-based type
// used throughout the logging system to indicate the importance or criticality
// of a message. The defined levels, in order of increasing verbosity, are:
// LevelFatal, LevelSilent, LevelError, LevelInfo, LevelWarn, and LevelDebug.
type Level int

// MarshalText implements the encoding.TextMarshaler interface to convert a Level
// to its text representation for serialization (e.g., JSON or YAML). It returns
// the string representation of the Level as a byte slice.
//
// Returns:
//   - bytes ([]byte): The string representation of the Level as a byte slice.
//   - err (error): Always nil, as marshaling a Level to text cannot fail.
func (l Level) MarshalText() (bytes []byte, err error) {
	bytes = []byte(l.String())

	return
}

// UnmarshalText implements the encoding.TextUnmarshaler interface to parse a
// text representation (e.g., from JSON or YAML) into a Level. It matches the input
// text against known level strings in the s array and sets the Level accordingly.
//
// Parameters:
//   - text ([]byte): The text representation of the Level to parse.
//
// Returns:
//   - err (error): Returns ErrUnknownLevel wrapped with the invalid text if the
//     input does not match any known level string; otherwise, nil.
func (l *Level) UnmarshalText(text []byte) (err error) {
	str := string(text)

	for i, v := range s {
		if v == str {
			*l = Level(i)

			return
		}
	}

	err = fmt.Errorf("%w (%s)", ErrUnknownLevel, s)

	return
}

// Int returns the integer value of the Level, allowing direct access to its
// underlying numeric representation for comparisons or indexing.
//
// Returns:
//   - level (int): The integer value of the Level.
func (l Level) Int() (level int) {
	level = int(l)

	return
}

// String returns the string representation of the Level, mapping its integer value
// to a lowercase label for use in log output or display. If the Level's integer
// value is out of range (i.e., less than 0 or greater than or equal to the length
// of the s array), it returns "unknown".
//
// Returns:
//   - level (string): The string representation of the Level, or "unknown" if invalid.
func (l Level) String() (level string) {
	if l.Int() < 0 || l.Int() >= len(s) {
		level = "unknown"

		return
	}

	level = s[l.Int()]

	return
}

// IsValid checks whether the Level has a valid integer value that corresponds to
// one of the defined logging levels (i.e., within the range of the s array).
//
// Returns:
//   - valid (bool): True if the Level is valid (between 0 and len(s)-1), false otherwise.
func (l Level) IsValid() (valid bool) {
	valid = l.Int() >= 0 && l.Int() < len(s)

	return
}

const (
	// LevelFatal represents critical errors that may cause the application to
	// terminate. Use this for unrecoverable conditions, such as failure to connect
	// to a critical service or data corruption. It has the highest severity (lowest
	// integer value).
	LevelFatal Level = iota
	// LevelSilent suppresses all logging output. When set as the logger's level,
	// no messages are emitted, regardless of their severity. Use this in production
	// environments to disable logging or minimize output.
	LevelSilent
	// LevelError indicates errors that require immediate attention but do not halt
	// program execution. Examples include failed API calls, invalid user input, or
	// resource unavailability.
	LevelError
	// LevelInfo captures informational messages about normal application operation,
	// such as successful initialization, user actions, or system state changes.
	LevelInfo
	// LevelWarn denotes warnings for potential issues or unexpected conditions that
	// do not prevent normal operation but may warrant investigation. Examples include
	// deprecated API usage or resource usage nearing limits.
	LevelWarn
	// LevelDebug provides detailed context for troubleshooting and development. Use
	// this for verbose output, such as variable states, function call traces, or
	// detailed system diagnostics, typically enabled in development or debugging
	// environments.
	LevelDebug
)

// s maps Level values to their string representations. It is used by the String()
// method to convert a Level to its corresponding lowercase label. The array is
// indexed by the integer value of the Level, with indices 0 to 5 corresponding to
// LevelFatal through LevelDebug. Out-of-range indices are handled safely by String()
// to return "unknown".
var s = [...]string{"fatal", "silent", "error", "info", "warn", "debug"}

var (
	// ErrUnknownLevel is an error returned when an invalid or unrecognized level string
	// is provided during unmarshaling or other operations that require a valid Level.
	ErrUnknownLevel = errors.New("unknown level")
)
