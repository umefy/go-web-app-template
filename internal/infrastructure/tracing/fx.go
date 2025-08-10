package tracing

import (
	"github.com/umefy/go-web-app-template/internal/infrastructure/tracing/opentelemetry"
	"go.uber.org/fx"
)

var Module = fx.Module("tracing",
	fx.Provide(opentelemetry.NewTracerProvider),
)
