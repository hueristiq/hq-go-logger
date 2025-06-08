# hq-go-logger

![made with go](https://img.shields.io/badge/made%20with-Go-1E90FF.svg) [![go report card](https://goreportcard.com/badge/github.com/hueristiq/hq-go-logger)](https://goreportcard.com/report/github.com/hueristiq/hq-go-logger) [![license](https://img.shields.io/badge/license-MIT-gray.svg?color=1E90FF)](https://github.com/hueristiq/hq-go-logger/blob/master/LICENSE) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-1E90FF.svg) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/hq-go-logger.svg?style=flat&color=1E90FF)](https://github.com/hueristiq/hq-go-logger/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/hq-go-logger.svg?style=flat&color=1E90FF)](https://github.com/hueristiq/hq-go-logger/issues?q=is:issue+is:closed) [![contribution](https://img.shields.io/badge/contributions-welcome-1E90FF.svg)](https://github.com/hueristiq/hq-go-logger/blob/master/CONTRIBUTING.md)

`hq-go-logger` is a [Go (Golang)](https://golang.org/) package for structured logging.

## Resources

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
	- [Basic Usage with DefaultLogger](#basic-usage-with-defaultlogger)
	- [Custom Logger Configuration](#custom-logger-configuration)
- [Contributing](#contributing)
- [Licensing](#licensing)

## Features

- **Structured Logging:** Attach metadata (key-value pairs) to log messages for enhanced context, such as request IDs or system metrics.
- **Multiple Log Levels:** Supports six levels (`Fatal`, `Silent`, `Error`, `Info`, `Warn`, `Debug`) for categorizing message severity.
- **Custom Formatters:** Swap or extend formatters to produce output in various formats (e.g., colorized console output, JSON, Logfmt).
- **Flexible Writers:** Route logs to multiple destinations, such as console, files, or external logging services.
- **Thread-Safe:** Ensures safe concurrent logging with thread-safe formatter and writer implementations.
- **Extensible API:** Provides clear interfaces (`Formatter`, `Writer`) for custom implementations and integration with existing systems.

## Installation

To install `hq-go-logger`, run the following command in your Go project:

```bash
go get -v -u github.com/hueristiq/hq-go-logger
```

Make sure your Go environment is set up properly (Go 1.18 or later is recommended).

## Usage

Below are examples demonstrating how to use the different features of the `hq-go-logger` package.

### Basic Usage with DefaultLogger

```go
package main

import (
	hqgologger "github.com/hueristiq/hq-go-logger"
	"github.com/hueristiq/hq-go-logger/formatter"
	"github.com/hueristiq/hq-go-logger/levels"
)

func main() {
	hqgologger.DefaultLogger.SetLevel(levels.LevelDebug)
	hqgologger.DefaultLogger.SetFormatter(formatter.NewConsoleFormatter(&formatter.ConsoleFormatterConfiguration{
		Colorize: true,
	}))

	hqgologger.Print("Print message", hqgologger.WithLabel("PRINT"), hqgologger.WithString("app", "default"))
	hqgologger.Info("Info message", hqgologger.WithLabel("INFO"), hqgologger.WithString("app", "default"))
	hqgologger.Warn("Warn message", hqgologger.WithLabel("WARN"), hqgologger.WithString("app", "default"))
	hqgologger.Error("Error message", hqgologger.WithLabel("ERROR"), hqgologger.WithString("app", "default"))
	hqgologger.Fatal("Fatal message", hqgologger.WithLabel("FATAL"), hqgologger.WithString("app", "default"))
}
```

### Custom Logger Configuration

```go
package main

import (
	hqgologger "github.com/hueristiq/hq-go-logger"
	"github.com/hueristiq/hq-go-logger/formatter"
	"github.com/hueristiq/hq-go-logger/levels"
	"github.com/hueristiq/hq-go-logger/writer"
)

func main() {
	logger := hqgologger.NewLogger()

	logger.SetLevel(levels.LevelDebug)
	logger.SetFormatter(formatter.NewConsoleFormatter(&formatter.ConsoleFormatterConfiguration{
		Colorize: true,
	}))
	logger.SetWriter(writer.NewConsoleWriter())

	logger.Print("Print message", hqgologger.WithLabel("PRINT"), hqgologger.WithString("app", "new"))
	logger.Info("Info message", hqgologger.WithLabel("INFO"), hqgologger.WithString("app", "new"))
	logger.Warn("Warn message", hqgologger.WithLabel("WARN"), hqgologger.WithString("app", "new"))
	logger.Error("Error message", hqgologger.WithLabel("ERROR"), hqgologger.WithString("app", "new"))
	logger.Fatal("Fatal message", hqgologger.WithLabel("FATAL"), hqgologger.WithString("app", "new"))
}
```

## Contributing

Contributions are welcome and encouraged! Feel free to submit [Pull Requests](https://github.com/hueristiq/hq-go-logger/pulls) or report [Issues](https://github.com/hueristiq/hq-go-logger/issues). For more details, check out the [contribution guidelines](https://github.com/hueristiq/hq-go-logger/blob/master/CONTRIBUTING.md).

A big thank you to all the [contributors](https://github.com/hueristiq/hq-go-logger/graphs/contributors) for your ongoing support!

![contributors](https://contrib.rocks/image?repo=hueristiq/hq-go-logger&max=500)

## Licensing

This package is licensed under the [MIT license](https://opensource.org/license/mit). You are free to use, modify, and distribute it, as long as you follow the terms of the license. You can find the full license text in the repository - [Full MIT license text](https://github.com/hueristiq/hq-go-logger/blob/master/LICENSE).