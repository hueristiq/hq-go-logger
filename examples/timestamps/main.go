package main

import (
	hqgologger "github.com/hueristiq/hq-go-logger"
)

func main() {
	hqgologger.Print("Print message with timestamp", hqgologger.WithLabel("PRINT"))
	hqgologger.Print("Print message without timestamp", hqgologger.WithLabel("PRINT"), hqgologger.WithoutTimestamp())
}
