package graphql

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"path"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/umefy/go-web-app-template/internal/app"
	"github.com/umefy/go-web-app-template/internal/domain/config"
	domainError "github.com/umefy/go-web-app-template/internal/domain/error"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router/middleware"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func NewGraphqlRouter(app *app.App) http.Handler {
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

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return checkWsOrigin(app, r)
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
			app.LoggerService.InfoContext(ctx, "WebSocket established",
				slog.String("request_id", middleware.GetReqID(ctx)),
			)
			return ctx, &initPayload, nil
		},
		ErrorFunc: func(ctx context.Context, err error) {
			var wsErr transport.WebsocketError
			if errors.As(err, &wsErr) {
				switch {
				case websocket.IsCloseError(wsErr.Err, websocket.CloseNormalClosure, websocket.CloseGoingAway):
					return
				case errors.Is(wsErr.Err, websocket.ErrCloseSent):
					return
				}
			}
			app.LoggerService.ErrorContext(ctx, "WebSocket error",
				slog.String("request_id", middleware.GetReqID(ctx)),
				slog.Any("error", err),
			)
		},
		CloseFunc: func(ctx context.Context, closeCode int) {
			app.LoggerService.InfoContext(ctx, "WebSocket closed",
				slog.String("request_id", middleware.GetReqID(ctx)),
				slog.Int("close_code", closeCode),
			)
		},
	})

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
	}

	// Create a router that handles WebSocket connections properly
	r := router.NewRouter()

	// Handle playground in development
	if appEnv == config.AppEnvDev {
		r.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	}

	// Handle GraphQL requests
	r.Handle("/", srv)

	return r
}

func checkWsOrigin(app *app.App, r *http.Request) bool {
	origin := r.Header.Get("Origin")
	if origin == "" {
		return false
	}

	for _, allowedOrigin := range app.ConfigService.GetHttpServerConfig().AllowedOrigins {
		// "*" matches anything
		if allowedOrigin == "*" {
			return true
		}

		// path.Match requires full match, use it as a simple glob
		ok, err := path.Match(allowedOrigin, origin)
		if err == nil && ok {
			return true
		}
	}

	return false
}
