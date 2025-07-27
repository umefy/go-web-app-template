package app

import (
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	"github.com/umefy/go-web-app-template/internal/core/config"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repo"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	greeterSvc "github.com/umefy/go-web-app-template/internal/service/greeter"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
)

type App struct {
	Arguments      config.Options
	Config         config.Config
	Logger         logger.Logger
	UserService    userSvc.Service
	UserRepository userRepo.Repository
	GreeterService greeterSvc.Service
	DbQuery        *query.Query
}

func New(configOptions config.Options) (*App, error) {
	app, err := InitializeApp(configOptions)
	if err != nil {
		return nil, err
	}
	return app, nil
}
