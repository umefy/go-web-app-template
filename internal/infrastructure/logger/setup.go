package logger

import (
	"log/slog"
	"os"

	"github.com/umefy/go-web-app-template/internal/core/config"
	"github.com/umefy/godash/logger"
)

func NewLogger(config config.Config) Logger {
	logLevel := GetLogLevel(config)

	loggerOpts := logger.NewLoggerOps(true, os.Stdout, logLevel, true, "source", 4)
	logger := logger.New(loggerOpts, func(handler slog.Handler) slog.Handler {
		return handler.WithAttrs([]slog.Attr{
			slog.Int("pid", os.Getpid()),
		})
	})

	return NewAppLogger(logger)
}

func GetLogLevel(config config.Config) slog.Level {
	loggingConfig := config.GetLoggingConfig()

	switch loggingConfig.Level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
