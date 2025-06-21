package middleware

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/umefy/godash/logger"
)

func Logger(logger *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ctx := r.Context()
			ctx = logger.WithValue(ctx, slog.String("request_id", getReqID(ctx)))

			// We don't want to log the source file and line number from a pkg folder
			loggerHandler := logger.GetHandler()
			loggerHandler.AddSource = false
			loggerWithoutSource := slog.New(&loggerHandler)

			loggerWithoutSource.InfoContext(ctx, "HTTP Request start",
				slog.String("request_id", getReqID(ctx)),
				slog.String("method", r.Method),
				slog.String("uri", r.RequestURI),
				slog.String("content_type", r.Header.Get("Content-Type")),
				slog.String("host", r.Host),
				slog.String("remote_ip", ExtractIP(r)))

			// Wrap ResponseWriter to capture status code
			ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(ww, r.WithContext(ctx))
			duration := time.Since(start)

			loggerWithoutSource.InfoContext(ctx,
				"HTTP Request done",
				slog.String("request_id", getReqID(ctx)),
				slog.String("method", r.Method),
				slog.String("uri", r.RequestURI),
				slog.String("content_type", r.Header.Get("Content-Type")),
				slog.Int("status", ww.statusCode),
				slog.String("host", r.Host),
				slog.String("remote_ip", ExtractIP(r)),
				slog.Duration("latency", duration),
			)
		})
	}
}

func getReqID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}

// ResponseWriter wrapper to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// ExtractIP extracts the real client IP, handling proxies
func ExtractIP(r *http.Request) string {
	// Check for a forwarded IP (common in reverse proxies)
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0]) // Take the first IP
	}

	// Fallback to direct RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // Use full RemoteAddr if parsing fails
	}

	return host
}
