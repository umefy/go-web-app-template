package tracing

import (
	"github.com/google/wire"
	"github.com/umefy/go-web-app-template/internal/infrastructure/tracing/opentelemetry"
)

var WireSet = wire.NewSet(
	opentelemetry.NewTraceProvider,
)
