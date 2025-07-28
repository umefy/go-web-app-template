package logger

import (
	"context"
	"log/slog"

	"github.com/umefy/godash/logger"
)

type Logger interface {
	Info(msg string, args ...slog.Attr)
	Error(msg string, args ...slog.Attr)
	Debug(msg string, args ...slog.Attr)
	Warn(msg string, args ...slog.Attr)

	InfoContext(ctx context.Context, msg string, args ...slog.Attr)
	ErrorContext(ctx context.Context, msg string, args ...slog.Attr)
	DebugContext(ctx context.Context, msg string, args ...slog.Attr)
	WarnContext(ctx context.Context, msg string, args ...slog.Attr)

	GetLogger() *logger.Logger
}

type appLogger struct {
	logger *logger.Logger
}

var _ Logger = (*appLogger)(nil)

func NewAppLogger(logger *logger.Logger) *appLogger {
	return &appLogger{logger: logger}
}

func (l *appLogger) Info(msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.Info(msg, attrs...)
}
func (l *appLogger) InfoContext(ctx context.Context, msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.InfoContext(ctx, msg, attrs...)
}

func (l *appLogger) Error(msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.Error(msg, attrs...)
}

func (l *appLogger) ErrorContext(ctx context.Context, msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.ErrorContext(ctx, msg, attrs...)
}

func (l *appLogger) Debug(msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.Debug(msg, attrs...)
}

func (l *appLogger) DebugContext(ctx context.Context, msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.DebugContext(ctx, msg, attrs...)
}

func (l *appLogger) Warn(msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.Warn(msg, attrs...)
}

func (l *appLogger) WarnContext(ctx context.Context, msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.WarnContext(ctx, msg, attrs...)
}

func (l *appLogger) GetLogger() *logger.Logger {
	return l.logger
}
