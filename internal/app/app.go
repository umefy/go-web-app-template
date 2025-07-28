package app

import (
	"github.com/umefy/go-web-app-template/internal/core/config"
	orderRepo "github.com/umefy/go-web-app-template/internal/domain/order/repo"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repo"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	greeterSvc "github.com/umefy/go-web-app-template/internal/service/greeter"
	orderSvc "github.com/umefy/go-web-app-template/internal/service/order"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
)

type App struct {
	Arguments       config.Options
	Config          config.Config
	Logger          logger.Logger
	UserService     userSvc.Service
	UserRepository  userRepo.Repository
	OrderService    orderSvc.Service
	OrderRepository orderRepo.Repository
	GreeterService  greeterSvc.Service
	DbQuery         *database.Query
}

func New(configOptions config.Options) (*App, error) {
	app, err := InitializeApp(configOptions)
	if err != nil {
		return nil, err
	}
	return app, nil
}
