package app

import (
	"fmt"

	configSvc "github.com/umefy/go-web-app-template/app/config/service"
	loggerSvc "github.com/umefy/go-web-app-template/app/logger/service"
	userRepo "github.com/umefy/go-web-app-template/app/user/repository"
	userSvc "github.com/umefy/go-web-app-template/app/user/service"
	db "github.com/umefy/go-web-app-template/pkg/db/gormdb"
	"github.com/umefy/go-web-app-template/pkg/validation"
	"github.com/umefy/godash/logger"
)

type App struct {
	Arguments      Arguments
	Logger         *logger.Logger
	LoggerService  loggerSvc.Service
	UserService    userSvc.Service
	UserRepository userRepo.Repository
	ConfigService  configSvc.Service
	DB             *db.DB
}

type Arguments struct {
	Env        string
	ConfigPath string
}

var _ validation.Validate = (*Arguments)(nil)

func (a *Arguments) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Env, validation.In("dev", "test", "prod").Error("can only be set to dev, test or prod")),
	)
}

func New(args Arguments) (*App, error) {
	err := args.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalidate arguments error: %w", err)
	}
	app, err := InitializeApp(args)
	if err != nil {
		return nil, err
	}
	return app, nil
}
