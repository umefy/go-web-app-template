package greeter

import "go.uber.org/fx"

var Module = fx.Module("greeter",
	fx.Provide(fx.Annotate(
		NewService,
		fx.As(new(Service)),
	)),
)
