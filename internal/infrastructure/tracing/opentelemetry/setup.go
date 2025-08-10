package opentelemetry

import (
	"context"

	"github.com/umefy/go-web-app-template/internal/core/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

func NewTraceProvider(ctx context.Context, config config.Config) (trace.TracerProvider, error) {
	traceConfig := config.GetTracingConfig()

	if !traceConfig.Enabled {
		return noop.NewTracerProvider(), nil
	}

	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(traceConfig.JaegerEndpoint), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(traceConfig.ServiceName),
			semconv.ServiceVersionKey.String(traceConfig.ServiceVersion),
		),
	)
	if err != nil {
		return nil, err
	}

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return traceProvider, nil
}
