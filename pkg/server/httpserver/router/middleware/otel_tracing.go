package middleware

import (
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func OTelTracing(tracerName string, tp trace.TracerProvider) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))

			tr := tp.Tracer(tracerName)
			ctx, span := tr.Start(ctx, fmt.Sprintf("%s %s", r.Method, r.URL.Path), trace.WithSpanKind(trace.SpanKindServer))
			span.SetAttributes(attribute.String("request_id", GetReqID(ctx)))
			defer span.End()

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
