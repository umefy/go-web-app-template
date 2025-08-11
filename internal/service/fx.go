package service

import (
	greeterSvc "github.com/umefy/go-web-app-template/internal/service/greeter"
	orderSvc "github.com/umefy/go-web-app-template/internal/service/order"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
	"go.uber.org/fx"
)

var Module = fx.Module("service",
	fx.Provide(
		fx.Annotate(
			userSvc.NewService,
			fx.As(new(userSvc.Service)),
		),
		fx.Annotate(
			orderSvc.NewService,
			fx.As(new(orderSvc.Service)),
		),
		fx.Annotate(
			greeterSvc.NewService,
			fx.As(new(greeterSvc.Service)),
		),
	),
)
