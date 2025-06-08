package levels

// Level represents the severity of a log message. It is an integer-based type
// used throughout the logging system to indicate the importance or criticality
// of a message.
type Level int

// Int returns the integer value of the Level.
//
// Returns:
//   - level (int): The integer value of the Level.
func (l Level) Int() (level int) {
	level = int(l)

	return
}

// String returns the string representation of the Level, mapping its integer value
// to a lowercase label for use in log output or display.
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

const (
	// LevelFatal represents critical errors that may cause the application to
	// terminate. Use this for unrecoverable conditions, such as failure to connect
	// to a critical service or data corruption.
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
