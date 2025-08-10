package user

import "go.uber.org/fx"

var Module = fx.Module("user",
	fx.Provide(fx.Annotate(
		NewService,
		fx.As(new(Service)),
	)),
)
