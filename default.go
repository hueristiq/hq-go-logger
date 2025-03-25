package logger

import (
	"go.source.hueristiq.com/logger/formatter"
	"go.source.hueristiq.com/logger/levels"
	"go.source.hueristiq.com/logger/writer"
)

var DefaultLogger *Logger

func init() {
	DefaultLogger = &Logger{}

	DefaultLogger.SetFormatter(formatter.NewConsoleFormatter(&formatter.ConsoleFormatterConfiguration{
		Colorize: true,
	}))
	DefaultLogger.SetWriter(writer.NewConsoleWriter())
	DefaultLogger.SetMaxLogLevel(levels.LevelDebug)
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
