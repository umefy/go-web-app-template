package graphql

import (
	"net/http"

	"go.uber.org/fx"
)

const (
	FX_TAG_NAME_GRAPHQL_ROUTER = `name:"graphqlRouter"`
)

var Module = fx.Module("graphqlRouter",
	fx.Provide(
		NewResolver,
		fx.Annotate(
			NewGraphqlRouter,
			fx.As(new(http.Handler)),
			fx.ResultTags(FX_TAG_NAME_GRAPHQL_ROUTER),
		),
	),
)
