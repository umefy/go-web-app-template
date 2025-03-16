package app

import (
	"log/slog"
	"os"

	configSvc "github.com/umefy/go-web-app-template/app/config/service"
	"github.com/umefy/godash/logger"
)

func newLogger(configSvc configSvc.Service) *logger.Logger {
	logLevel := getLogLevel(configSvc)

	loggerOpts := logger.NewLoggerOps(true, os.Stdout, logLevel, true, "source", 4)
	logger := logger.New(loggerOpts, func(handler slog.Handler) slog.Handler {
		return handler.WithAttrs([]slog.Attr{
			slog.Int("pid", os.Getpid()),
		})
	})

	return logger
}

func getLogLevel(configSvc configSvc.Service) slog.Level {
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
