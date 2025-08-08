package main

import (
	hqgoerrors "github.com/hueristiq/hq-go-errors"
	hqgologger "github.com/hueristiq/hq-go-logger"
)

func main() {
	err := hqgoerrors.New("root error example!", hqgoerrors.WithType("ERROR_TYPE"), hqgoerrors.WithField("FIELD_KEY_1", "FIELD_VALUE_1"), hqgoerrors.WithField("FIELD_KEY_2", "FIELD_VALUE_2"))

	hqgologger.Error("Error message", hqgologger.WithString("string-key", "string-value"), hqgologger.WithValue("value-key", "value-value"), hqgologger.WithError(err))
}
