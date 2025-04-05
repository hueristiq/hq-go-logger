package writer

import "github.com/hueristiq/hq-go-logger/levels"

// Writer is an abstraction for writing log data.
// It allows log messages to be sent to different output targets, such as the console,
// a file, or other custom destinations. The interface requires implementing a single method,
// Write, which accepts formatted log data and the corresponding log level.
//
// The log level parameter can be used by implementations to route or format the output
// differently based on the severity of the log message.
type Writer interface {
	Write(data []byte, level levels.Level)
}
