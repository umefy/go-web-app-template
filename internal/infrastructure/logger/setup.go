package logger

import (
	"log/slog"
	"os"

	configSvc "github.com/umefy/go-web-app-template/internal/domain/config/service"
	"github.com/umefy/godash/logger"
)

func NewLogger(configSvc configSvc.Service) *logger.Logger {
	logLevel := GetLogLevel(configSvc)

	loggerOpts := logger.NewLoggerOps(true, os.Stdout, logLevel, true, "source", 4)
	logger := logger.New(loggerOpts, func(handler slog.Handler) slog.Handler {
		return handler.WithAttrs([]slog.Attr{
			slog.Int("pid", os.Getpid()),
		})
	})

	return logger
}

func GetLogLevel(configSvc configSvc.Service) slog.Level {
	loggingConfig := configSvc.GetLoggingConfig()

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
