package main

import (
	hqgologger "github.com/hueristiq/hq-go-logger"
	hqgologgerformatter "github.com/hueristiq/hq-go-logger/formatter"
	hqgologgercolorizer "github.com/hueristiq/hq-go-logger/formatter/colorizer"
	hqgologgerlevels "github.com/hueristiq/hq-go-logger/levels"
	hqgologgerwriter "github.com/hueristiq/hq-go-logger/writer"
)

func main() {
	logger := hqgologger.NewLogger()

	logger.SetLevel(hqgologgerlevels.LevelDebug)

	fcfg := hqgologgerformatter.DefaultConsoleConfig()

	fcfg.Colorizer = hqgologgercolorizer.NewFatihColorizer()

	logger.SetFormatter(hqgologgerformatter.NewConsoleFormatter(fcfg))
	logger.SetWriter(hqgologgerwriter.NewConsoleWriter(hqgologgerwriter.DefaultConsoleWriterConfig()))

	logger.Print("Print message", hqgologger.WithLabel("PRINT"), hqgologger.WithString("string-key", "string-value"), hqgologger.WithValue("value-key", "value-value"))
	logger.Info("Info message", hqgologger.WithLabel("INFO"), hqgologger.WithString("string-key", "string-value"), hqgologger.WithValue("value-key", "value-value"))
	logger.Warn("Warn message", hqgologger.WithLabel("WARN"), hqgologger.WithString("string-key", "string-value"), hqgologger.WithValue("value-key", "value-value"))
	logger.Error("Error message", hqgologger.WithLabel("ERROR"), hqgologger.WithString("string-key", "string-value"), hqgologger.WithValue("value-key", "value-value"))
	logger.Fatal("Fatal message", hqgologger.WithLabel("FATAL"), hqgologger.WithString("string-key", "string-value"), hqgologger.WithValue("value-key", "value-value"))
}
