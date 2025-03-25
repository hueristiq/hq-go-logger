# hq-go-logger

![made with go](https://img.shields.io/badge/made%20with-Go-1E90FF.svg) [![go report card](https://goreportcard.com/badge/github.com/hueristiq/hq-go-logger)](https://goreportcard.com/report/github.com/hueristiq/hq-go-logger) [![license](https://img.shields.io/badge/license-MIT-gray.svg?color=1E90FF)](https://github.com/hueristiq/hq-go-logger/blob/master/LICENSE) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-1E90FF.svg) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/hq-go-logger.svg?style=flat&color=1E90FF)](https://github.com/hueristiq/hq-go-logger/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/hq-go-logger.svg?style=flat&color=1E90FF)](https://github.com/hueristiq/hq-go-logger/issues?q=is:issue+is:closed) [![contribution](https://img.shields.io/badge/contributions-welcome-1E90FF.svg)](https://github.com/hueristiq/hq-go-logger/blob/master/CONTRIBUTING.md)

`hq-go-logger` is a [Go(Golang)](https://golang.org/) package for structured logging. It provides a flexible API for logging messages with varying severity levels, custom formatting, and output destinations.

## Resources

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [Licensing](#licensing)

## Features

- **Structured Logging:** Log messages with associated metadata for richer context.
- **Multiple Log Levels:** Built-in support for levels such as Fatal, Error, Info, Warn, Debug, and Silent.
- **Custom Formatters:** Easily swap out or extend log message formatting to suit your needs (e.g., colorized console output).
- **Flexible Writers:** Route log output to various destinations like the console, files, or external logging services.
- **Thread-Safe:** Designed with concurrency in mind for safe use in multi-goroutine applications.
- **Extensible API:** Provides clear interfaces (Formatter, Writer) that allow you to extend functionality or integrate with existing systems.

## Installation

To install `hq-go-logger`, run:

```bash
go get -v -u go.source.hueristiq.com/logger
```

Make sure your Go environment is set up properly (Go 1.x or later is recommended).

## Usage

Below are examples demonstrating how to use the different features of the `hq-go-logger` package.

```go
package main

import (
	"go.source.hueristiq.com/logger"
	"go.source.hueristiq.com/logger/formatter"
	"go.source.hueristiq.com/logger/levels"
)

func main() {
	logger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	logger.DefaultLogger.SetFormatter(formatter.NewCLI(&formatter.CLIOptions{
		Colorize: true,
	}))


	logger.Print().Msg("Print message")
	logger.Info().Msg("Info message")
	logger.Warn().Msg("Warn message")
	logger.Error().Msg("Error message")
	logger.Fatal().Msg("Fatal message")
}
```

## Contributing

Contributions are welcome and encouraged! Feel free to submit [Pull Requests](https://github.com/hueristiq/hq-go-logger/pulls) or report [Issues](https://github.com/hueristiq/hq-go-logger/issues). For more details, check out the [contribution guidelines](https://github.com/hueristiq/hq-go-logger/blob/master/CONTRIBUTING.md).

A big thank you to all the [contributors](https://github.com/hueristiq/hq-go-logger/graphs/contributors) for your ongoing support!

![contributors](https://contrib.rocks/image?repo=hueristiq/hq-go-logger&max=500)

## Licensing

This package is licensed under the [MIT license](https://opensource.org/license/mit). You are free to use, modify, and distribute it, as long as you follow the terms of the license. You can find the full license text in the repository - [Full MIT license text](https://github.com/hueristiq/hq-go-logger/blob/master/LICENSE).