package main

import (
	hqgologger "github.com/hueristiq/hq-go-logger"
)

func main() {
	hqgologger.Print("Print message")
	hqgologger.Info("Info message")
	hqgologger.Warn("Warn message")
	hqgologger.Error("Error message")
	hqgologger.Fatal("Fatal message")
}
