//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/umefy/go-web-app-template/internal/core/config"
	greeterSvc "github.com/umefy/go-web-app-template/internal/domain/greeter/service"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repository"
	userSvc "github.com/umefy/go-web-app-template/internal/domain/user/service"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
)

var WireSet = wire.NewSet(
	logger.NewLogger,
	database.WireSet,
	config.NewConfig,
	userSvc.WireSet,
	greeterSvc.WireSet,
	wire.Bind(new(userRepo.Repository), new(*database.UserRepository)),
	wire.Struct(new(App), "*"),
)

func InitializeApp(configOptions config.Options) (*App, error) {
	wire.Build(WireSet)
	return &App{}, nil
}
