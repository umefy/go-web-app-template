package graphql

import (
	"context"
	"errors"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/umefy/go-web-app-template/internal/app"
	"github.com/umefy/go-web-app-template/internal/domain/config"
	domainError "github.com/umefy/go-web-app-template/internal/domain/error"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func NewGraphqlRouter(app *app.App) http.Handler {
	r := router.NewRouter()

	graphqlConfig := Config{
		Resolvers: &Resolver{
			UserService:   app.UserService,
			LoggerService: app.LoggerService,
			DbQuery:       app.DbQuery,
		},
	}

	configSvc := app.ConfigService
	appEnv := configSvc.GetAppConfig().Env

	srv := handler.New(NewExecutableSchema(graphqlConfig))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		var domainErr *domainError.Error
		if errors.As(e, &domainErr) {
			err.Message = domainErr.Message
			err.Extensions = map[string]any{"code": domainErr.Code}
		}

		return err
	})

	if appEnv == config.AppEnvDev {
		srv.Use(extension.Introspection{})
		r.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	}

	r.Handle("/", srv)

	return r
}
