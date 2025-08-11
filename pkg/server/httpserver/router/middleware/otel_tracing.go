package middleware

import (
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func OTelTracing(tracerName string, tp trace.TracerProvider) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Check if this is a WebSocket upgrade request
			if IsWebSocketUpgrade(r) {
				// For WebSocket connections, don't wrap the response writer
				next.ServeHTTP(w, r)
				return
			}

			ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))

			tr := tp.Tracer(tracerName)
			ctx, span := tr.Start(ctx, fmt.Sprintf("%s %s", r.Method, r.URL.Path), trace.WithSpanKind(trace.SpanKindServer))
			span.SetAttributes(attribute.String("request_id", GetReqID(ctx)))
			defer span.End()

			// Wrap ResponseWriter to capture status code for regular HTTP requests
			ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(ww, r.WithContext(ctx))

			span.SetAttributes(attribute.Int("status_code", ww.statusCode))

			if ww.statusCode >= 400 {
				span.SetStatus(codes.Error, http.StatusText(ww.statusCode))
			}
		})
	}
}
