package service

import (
	"context"
	"log/slog"
)

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
	Warn(msg string, args ...any)

	InfoContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
}

type Service interface {
	Info(msg string, args ...slog.Attr)
	Error(msg string, args ...slog.Attr)
	Debug(msg string, args ...slog.Attr)
	Warn(msg string, args ...slog.Attr)

	InfoContext(ctx context.Context, msg string, args ...slog.Attr)
	ErrorContext(ctx context.Context, msg string, args ...slog.Attr)
	DebugContext(ctx context.Context, msg string, args ...slog.Attr)
	WarnContext(ctx context.Context, msg string, args ...slog.Attr)
}

type loggerService struct {
	logger Logger
}

var _ Service = (*loggerService)(nil)

func NewService(logger Logger) *loggerService {
	return &loggerService{logger: logger}
}

func (l *loggerService) Info(msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.Info(msg, attrs...)
}
func (l *loggerService) InfoContext(ctx context.Context, msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.InfoContext(ctx, msg, attrs...)
}

func (l *loggerService) Error(msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.Error(msg, attrs...)
}

func (l *loggerService) ErrorContext(ctx context.Context, msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.ErrorContext(ctx, msg, attrs...)
}

func (l *loggerService) Debug(msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.Debug(msg, attrs...)
}

func (l *loggerService) DebugContext(ctx context.Context, msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.DebugContext(ctx, msg, attrs...)
}

func (l *loggerService) Warn(msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.Warn(msg, attrs...)
}

func (l *loggerService) WarnContext(ctx context.Context, msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i, attr := range args {
		attrs[i] = attr
	}
	l.logger.WarnContext(ctx, msg, attrs...)
}
