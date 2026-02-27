package a

import (
	"log/slog"

	"go.uber.org/zap"
)

func test() {
	slog.Info("Starting server")    // want "must start with lowercase"
	slog.Info("запуск сервера")     // want "log message must be in English"
	slog.Info("server started!!!")  // want "log message must not contain special characters or emoji"
	slog.Info("user password: 123") // want "log message must not contain special characters or emoji" "log message contains potentially sensitive data"
	logger := zap.NewExample()
	logger.Info("Starting server")    // want "must start with lowercase"
	logger.Info("запуск сервера")     // want "log message must be in English"
	logger.Info("server started!!!")  // want "log message must not contain special characters or emoji"
	logger.Info("user password: 123") // want "log message must not contain special characters or emoji" "log message contains potentially sensitive data"
}
