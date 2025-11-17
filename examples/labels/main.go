package main

import (
	hqgologger "github.com/hueristiq/hq-go-logger"
)

func main() {
	hqgologger.Print("No Label", hqgologger.WithLabel(""))
	hqgologger.Print("Print message", hqgologger.WithLabel("PRINT"))
	hqgologger.Info("Info message", hqgologger.WithLabel("INFO"))
	hqgologger.Warn("Warn message", hqgologger.WithLabel("WARN"))
	hqgologger.Error("Error message", hqgologger.WithLabel("ERROR"))
	hqgologger.Fatal("Fatal message", hqgologger.WithLabel("FATAL"))
}
