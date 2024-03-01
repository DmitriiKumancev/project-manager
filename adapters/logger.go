package adapters

import (
	"log/slog"
	"os"
)

var defaultLogger *slog.Logger


func init() {
	defaultLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(defaultLogger)
}

func GetLogger() *slog.Logger {
	return defaultLogger
}

func SetLogger(logger *slog.Logger) {
	defaultLogger = logger
	slog.SetDefault(logger)
}