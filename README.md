# hq-go-logger

![made with go](https://img.shields.io/badge/made%20with-Go-1E90FF.svg) [![go report card](https://goreportcard.com/badge/github.com/hueristiq/hq-go-logger)](https://goreportcard.com/report/github.com/hueristiq/hq-go-logger) [![license](https://img.shields.io/badge/license-MIT-gray.svg?color=1E90FF)](https://github.com/hueristiq/hq-go-logger/blob/master/LICENSE) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-1E90FF.svg) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/hq-go-logger.svg?style=flat&color=1E90FF)](https://github.com/hueristiq/hq-go-logger/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/hq-go-logger.svg?style=flat&color=1E90FF)](https://github.com/hueristiq/hq-go-logger/issues?q=is:issue+is:closed) [![contribution](https://img.shields.io/badge/contributions-welcome-1E90FF.svg)](https://github.com/hueristiq/hq-go-logger/blob/master/CONTRIBUTING.md)

`hq-go-logger` is a [Go (Golang)](https://golang.org/) package for flexible and extensible structured logging. It provides a robust logging system with support for customizable log levels, formatters, writers, and colorized console output.

## Resources

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
	- [Basic Usage with `DefaultLogger`](#basic-usage-with-defaultlogger)
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

### Basic Usage with `DefaultLogger`

The `DefaultLogger` is pre-configured with a `LevelDebug` threshold, a `Console` formatter with colorized output, and a `Console` writer (stdout for `LevelSilent`, stderr for others). This example shows basic logging with metadata and custom labels.

```go
package main

import (
	hqgologger "github.com/hueristiq/hq-go-logger"
)

func main() {
	hqgologger.Print("Application started", hqgologger.WithLabel("START"), hqgologger.WithString("app", "my-app"))
	hqgologger.Info("Processing request", hqgologger.WithString("request_id", "12345"))
	hqgologger.Warn("Resource usage high", hqgologger.WithString("memory", "80%"))
	hqgologger.Error("Failed to connect", hqgologger.WithError(errors.New("connection timeout")))
	hqgologger.Fatal("Critical failure", hqgologger.WithLabel("CRIT"), hqgologger.WithString("app", "my-app"))
}
```

```
2025-08-08T13:45:00Z [START] Application started app=my-app
2025-08-08T13:45:00Z [INF] Processing request request_id=12345
2025-08-08T13:45:00Z [WRN] Resource usage high memory=80%
2025-08-08T13:45:00Z [ERR] Failed to connect error=connection timeout
2025-08-08T13:45:00Z [CRIT] Critical failure app=my-app
```

### Custom Logger Configuration

For more control, create a custom `Logger` instance with specific formatters and writers. This example configures a logger with a custom `Console` formatter and writer, disabling newlines for compact output.

```go
package main

import (
	"errors"

	hqgologger "github.com/hueristiq/hq-go-logger"
	hqgologgerformatter "github.com/hueristiq/hq-go-logger/formatter"
	hqgologgercolorizer "github.com/hueristiq/hq-go-logger/formatter/colorizer"
	hqgologgerlevels "github.com/hueristiq/hq-go-logger/levels"
	hqgologgerwritter "github.com/hueristiq/hq-go-logger/writer"
)

func main() {
	logger := hqgologger.NewLogger()

	logger.SetLevel(hqgologgerlevels.LevelDebug)
	logger.SetFormatter(hqgologgerformatter.NewConsoleFormatter(&hqgologgerformatter.ConsoleFormatterConfiguration{
		IncludeTimestamp: true,
		TimestampFormat:  "2006-01-02 15:04:05",
		IncludeLabel:     true,
		Colorize:         true,
		Colorizer:        hqgologgercolorizer.NewFatihColorizer(),
	}))
	logger.SetWriter(hqgologgerwritter.NewConsoleWriter(&hqgologgerwritter.ConsoleWriterConfiguration{
		ForceStdout:     true, // Route all logs to stdout
		DisableNewline:  true, // No newlines for compact output
	}))

	logger.Print("Application started", hqgologger.WithLabel("START"), hqgologger.WithString("app", "custom"))
	logger.Info("Processing request", hqgologger.WithString("request_id", "67890"))
	logger.Error("Connection failed", hqgologger.WithError(errors.New("network error")))
}
```

```
2025-08-08 13:45:05 [START] Application started app=custom 2025-08-08 13:45:05 [INF] Processing request request_id=67890 2025-08-08 13:45:05 [ERR] Connection failed error=network error
```

## Contributing

Contributions are welcome and encouraged! Feel free to submit [Pull Requests](https://github.com/hueristiq/hq-go-logger/pulls) or report [Issues](https://github.com/hueristiq/hq-go-logger/issues). For more details, check out the [contribution guidelines](https://github.com/hueristiq/hq-go-logger/blob/master/CONTRIBUTING.md).

A big thank you to all the [contributors](https://github.com/hueristiq/hq-go-logger/graphs/contributors) for your ongoing support!

![contributors](https://contrib.rocks/image?repo=hueristiq/hq-go-logger&max=500)

## Licensing

This package is licensed under the [MIT license](https://opensource.org/license/mit). You are free to use, modify, and distribute it, as long as you follow the terms of the license. You can find the full license text in the repository - [Full MIT license text](https://github.com/hueristiq/hq-go-logger/blob/master/LICENSE).