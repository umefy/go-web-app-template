package app

import (
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	configSvc "github.com/umefy/go-web-app-template/internal/domain/config/service"
	greeterSvc "github.com/umefy/go-web-app-template/internal/domain/greeter/service"
	loggerSvc "github.com/umefy/go-web-app-template/internal/domain/logger/service"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repository"
	userSvc "github.com/umefy/go-web-app-template/internal/domain/user/service"
	"github.com/umefy/go-web-app-template/internal/infrastructure/config"
	"github.com/umefy/godash/logger"
)

type App struct {
	Arguments      config.Options
	Logger         *logger.Logger
	LoggerService  loggerSvc.Service
	UserService    userSvc.Service
	UserRepository userRepo.Repository
	GreeterService greeterSvc.Service
	ConfigService  configSvc.Service
	DbQuery        *query.Query
}

func New(configOptions config.Options) (*App, error) {
	app, err := InitializeApp(configOptions)
	if err != nil {
		return nil, err
	}
	return app, nil
}
