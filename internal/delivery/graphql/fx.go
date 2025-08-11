package graphql

import (
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Module("graphqlRouter",
	fx.Provide(
		NewResolver,
		fx.Annotate(
			NewGraphqlRouter,
			fx.As(new(http.Handler)),
			fx.ResultTags(`name:"graphqlRouter"`),
		),
	),
)
