package greeter

import (
	"context"
	"log/slog"

	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	"go.opentelemetry.io/otel/trace"
)

type Service interface {
	SayHello(ctx context.Context, name string) (string, error)
}

type greetService struct {
	logger logger.Logger
	tp     trace.TracerProvider
}

var _ Service = (*greetService)(nil)

func NewService(logger logger.Logger, tp trace.TracerProvider) *greetService {
	return &greetService{logger: logger, tp: tp}
}

// SayHello implements Service.
func (g *greetService) SayHello(ctx context.Context, name string) (string, error) {
	tr := g.tp.Tracer("greeterService")
	_, span := tr.Start(ctx, "SayHello")
	defer span.End()

	g.logger.InfoContext(ctx, "SayHello", slog.String("name", name))
	return "Hello, " + name, nil
}
