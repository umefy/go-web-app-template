package graphql

import (
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
	"go.opentelemetry.io/otel/trace"
)

type Resolver struct {
	UserService    userSvc.Service
	Logger         logger.Logger
	TracerProvider trace.TracerProvider
}

func NewResolver(userService userSvc.Service, logger logger.Logger, tracerProvider trace.TracerProvider) *Resolver {
	return &Resolver{
		UserService:    userService,
		Logger:         logger,
		TracerProvider: tracerProvider,
	}
}
