package levels

// Level represents the severity of a log message.
// It is defined as an integer type and is used throughout the logging system
// to indicate the importance or criticality of a message.
//
// The valid Level values and their meanings are:
//
//   - LevelFatal  (0): Critical errors that may lead to application termination.
//   - LevelSilent (1): Logging is suppressed; no messages will be emitted.
//   - LevelError  (2): Error messages that require immediate attention but do not stop execution.
//   - LevelInfo   (3): Informational messages about normal application operation.
//   - LevelWarn   (4): Warning messages indicating potential issues or unexpected conditions.
//   - LevelDebug  (5): Debug-level messages providing detailed context for troubleshooting.
type Level int

// String returns the string representation of the Level.
// It maps the Level's integer value to a corresponding lowercase label.
//
// Supported mappings:
//
//   - 0 -> "fatal"
//   - 1 -> "silent"
//   - 2 -> "error"
//   - 3 -> "info"
//   - 4 -> "warn"
//   - 5 -> "debug"
//
// If the Level value is outside the range [0, 5], this function will panic due to an out-of-bounds slice index.
// To safely convert an arbitrary integer to a Level, validate it against the defined constants before calling String().
func (l Level) String() string {
	return [...]string{"fatal", "silent", "error", "info", "warn", "debug"}[l]
}

const (
	LevelFatal Level = iota
	LevelSilent
	LevelError
	LevelInfo
	LevelWarn
	LevelDebug
)
