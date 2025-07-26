//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/umefy/go-web-app-template/internal/core/config"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repository"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	greeterSvc "github.com/umefy/go-web-app-template/internal/service/greeter"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
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
