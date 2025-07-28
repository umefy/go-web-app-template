//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/umefy/go-web-app-template/internal/core/config"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	greeterSvc "github.com/umefy/go-web-app-template/internal/service/greeter"
	orderSvc "github.com/umefy/go-web-app-template/internal/service/order"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
)

var WireSet = wire.NewSet(
	logger.NewLogger,
	database.WireSet,
	config.NewConfig,
	userSvc.WireSet,
	orderSvc.WireSet,
	greeterSvc.WireSet,
	wire.Struct(new(App), "*"),
)

func InitializeApp(configOptions config.Options) (*App, error) {
	wire.Build(WireSet)
	return &App{}, nil
}
