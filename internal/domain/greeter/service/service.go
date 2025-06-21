package service

import (
	"context"
	"log/slog"

	loggerSrv "github.com/umefy/go-web-app-template/internal/domain/logger/service"
)

type Service interface {
	SayHello(ctx context.Context, name string) (string, error)
}

type greetService struct {
	loggerService loggerSrv.Service
}

var _ Service = (*greetService)(nil)

func NewService(loggerService loggerSrv.Service) *greetService {
	return &greetService{loggerService: loggerService}
}

// SayHello implements Service.
func (g *greetService) SayHello(ctx context.Context, name string) (string, error) {
	g.loggerService.InfoContext(ctx, "SayHello", slog.String("name", name))
	return "Hello, " + name, nil
}
