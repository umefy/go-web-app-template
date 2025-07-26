package service

import (
	"context"
	"log/slog"

	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
)

type Service interface {
	SayHello(ctx context.Context, name string) (string, error)
}

type greetService struct {
	logger logger.Logger
}

var _ Service = (*greetService)(nil)

func NewService(logger logger.Logger) *greetService {
	return &greetService{logger: logger}
}

// SayHello implements Service.
func (g *greetService) SayHello(ctx context.Context, name string) (string, error) {
	g.logger.InfoContext(ctx, "SayHello", slog.String("name", name))
	return "Hello, " + name, nil
}
