package app

import (
	"context"

	"github.com/umefy/go-web-app-template/internal/core/config"
	orderRepo "github.com/umefy/go-web-app-template/internal/domain/order/repo"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repo"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	greeterSvc "github.com/umefy/go-web-app-template/internal/service/greeter"
	orderSvc "github.com/umefy/go-web-app-template/internal/service/order"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	ctx             context.Context // context used for initial different infra
	Arguments       config.Options
	Config          config.Config
	Logger          logger.Logger
	UserService     userSvc.Service
	UserRepository  userRepo.Repository
	OrderService    orderSvc.Service
	OrderRepository orderRepo.Repository
	GreeterService  greeterSvc.Service
	DbQuery         *database.Query
	Tracer          trace.Tracer
}

func New(configOptions config.Options) (*App, error) {
	ctx := context.Background()
	app, err := InitializeApp(configOptions, ctx)
	if err != nil {
		return nil, err
	}
	return app, nil
}
