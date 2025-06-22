//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	configSvc "github.com/umefy/go-web-app-template/internal/domain/config/service"
	greeterSvc "github.com/umefy/go-web-app-template/internal/domain/greeter/service"
	loggerSvc "github.com/umefy/go-web-app-template/internal/domain/logger/service"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repository"
	userSvc "github.com/umefy/go-web-app-template/internal/domain/user/service"
	"github.com/umefy/go-web-app-template/internal/infrastructure/config"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	dashLogger "github.com/umefy/godash/logger"
)

var WireSet = wire.NewSet(
	logger.NewLogger,
	database.WireSet,
	wire.Bind(new(loggerSvc.Logger), new(*dashLogger.Logger)),
	config.LoadConfig,
	configSvc.WireSet,
	userSvc.WireSet,
	greeterSvc.WireSet,
	loggerSvc.WireSet,
	wire.Bind(new(userRepo.Repository), new(*database.UserRepository)),
	wire.Struct(new(App), "*"),
)

func InitializeApp(configOptions config.Options) (*App, error) {
	wire.Build(WireSet)
	return &App{}, nil
}
