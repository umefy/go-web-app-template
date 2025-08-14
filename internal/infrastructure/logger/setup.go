package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/umefy/go-web-app-template/internal/core/config"
	"github.com/umefy/godash/logger"
)

func NewLogger(config config.Config) Logger {
	loggingConfig := config.GetLoggingConfig()

	loggerOpts := logger.NewLoggerOps(
		loggingConfig.UseJson,
		getWriter(loggingConfig.Writer),
		getLogLevel(loggingConfig.Level),
		loggingConfig.AddSource,
		loggingConfig.SourceKey,
		4,
	)
	logger := logger.New(loggerOpts, func(handler slog.Handler) slog.Handler {
		return handler.WithAttrs([]slog.Attr{
			slog.Int("pid", os.Getpid()),
		})
	})

	return NewAppLogger(logger)
}

func getLogLevel(levelConfig string) slog.Level {
	switch strings.ToLower(levelConfig) {
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

func getWriter(writerConfig string) io.Writer {
	switch strings.ToLower(writerConfig) {
	case "stdout":
		return os.Stdout
	default:
		return os.Stdout
	}
}
