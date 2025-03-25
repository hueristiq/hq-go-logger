package levels

// Level represents the severity level of a log message.
// It is defined as an integer type and is used throughout the logging system
// to indicate the importance or criticality of a message.
type Level int

// String returns the string representation of the log level.
// It maps the Level's integer value to a corresponding string label.
// The mapping is defined as follows:
//
//	0 -> "fatal"
//	1 -> "silent"
//	2 -> "error"
//	3 -> "info"
//	4 -> "warn"
//	5 -> "debug"
//
// Note: This function uses a fixed array literal and indexes into it using the Level value.
// Ensure that the Level value is within the valid range to avoid a runtime panic.
func (l Level) String() string {
	return [...]string{"fatal", "silent", "error", "info", "warn", "debug"}[l]
}

// Predefined log level constants.
// The iota enumerator is used to assign consecutive integer values to these levels,
// starting from 0.
//
// LevelFatal:
//
//	Indicates the highest level of severity. It is used for critical errors that may
//	lead to the termination of the application.
//
// LevelSilent:
//
//	Indicates a level at which logging is suppressed or intentionally omitted.
//
// LevelError:
//
//	Used for error messages that highlight issues requiring immediate attention but do
//	not necessarily stop the program execution.
//
// LevelInfo:
//
//	Represents informational messages that provide general runtime information about
//	the application’s operation.
//
// LevelWarn:
//
//	Indicates warning messages that alert the user to potential issues or unexpected
//	conditions in the application.
//
// LevelDebug:
//
//	Used for debug-level messages that provide detailed context useful for developers.
const (
	LevelFatal Level = iota
	LevelSilent
	LevelError
	LevelInfo
	LevelWarn
	LevelDebug
)
