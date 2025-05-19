//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	configSvc "github.com/umefy/go-web-app-template/app/config/service"
	greeterSvc "github.com/umefy/go-web-app-template/app/greeter/service"
	loggerSvc "github.com/umefy/go-web-app-template/app/logger/service"
	userRepo "github.com/umefy/go-web-app-template/app/user/repository"
	userSvc "github.com/umefy/go-web-app-template/app/user/service"
	"github.com/umefy/godash/logger"
)

var WireSet = wire.NewSet(
	newLogger,
	newDB,
	wire.Bind(new(loggerSvc.Logger), new(*logger.Logger)),
	LoadConfig,
	configSvc.WireSet,
	userSvc.WireSet,
	greeterSvc.WireSet,
	loggerSvc.WireSet,
	userRepo.WireSet,
	wire.Struct(new(App), "*"),
)

func InitializeApp(args Arguments) (*App, error) {
	wire.Build(WireSet)
	return &App{}, nil
}
