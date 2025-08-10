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
	gqlgenExtension "github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/umefy/go-web-app-template/internal/core/config"
	"github.com/umefy/go-web-app-template/internal/delivery/errutil"
	"github.com/umefy/go-web-app-template/internal/delivery/graphql/dataloader"
	"github.com/umefy/go-web-app-template/internal/delivery/graphql/extension"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	orderSvc "github.com/umefy/go-web-app-template/internal/service/order"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router/middleware"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/fx"
)

type GraphqlRouterParams struct {
	fx.In

	UserService  userSvc.Service
	Logger       logger.Logger
	Config       config.Config
	DbQuery      *database.Query
	OrderService orderSvc.Service
}

func NewGraphqlRouter(params GraphqlRouterParams) http.Handler {
	graphqlConfig := Config{
		Resolvers: &Resolver{
			UserService: params.UserService,
			Logger:      params.Logger,
		},
	}

	appEnv := params.Config.GetEnv()

	srv := handler.New(NewExecutableSchema(graphqlConfig))

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return checkWsOrigin(params, r)
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
			params.Logger.InfoContext(ctx, "WebSocket established",
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
			params.Logger.ErrorContext(ctx, "WebSocket error",
				slog.String("request_id", middleware.GetReqID(ctx)),
				slog.Any("error", err),
			)
		},
		CloseFunc: func(ctx context.Context, closeCode int) {
			params.Logger.InfoContext(ctx, "WebSocket closed",
				slog.String("request_id", middleware.GetReqID(ctx)),
				slog.Int("close_code", closeCode),
			)
		},
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(gqlgenExtension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	srv.Use(&extension.TransactionExtension{
		DbQuery: params.DbQuery,
		Logger:  params.Logger,
	})

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		_, errMap := errutil.FormatError(e)
		err.Message = errMap["error"].(map[string]any)["message"].(string)
		err.Extensions = errMap

		return err
	})

	if appEnv == config.AppEnvDev {
		srv.Use(gqlgenExtension.Introspection{})
	}

	// Create a router that handles WebSocket connections properly
	r := router.NewRouter()

	dataloaderDeps := dataloader.LoaderDeps{
		OrderService: params.OrderService,
		Logger:       params.Logger,
	}
	// Handle playground in development
	if appEnv == config.AppEnvDev {
		r.Handle("/playground", dataloader.Middleware(playground.Handler("GraphQL playground", "/graphql"), dataloaderDeps))
	}

	// Handle GraphQL requests
	r.Handle("/", dataloader.Middleware(srv, dataloaderDeps))

	return r
}

func checkWsOrigin(params GraphqlRouterParams, r *http.Request) bool {
	origin := r.Header.Get("Origin")
	if origin == "" {
		return false
	}

	for _, allowedOrigin := range params.Config.GetHttpServerConfig().AllowedOrigins {
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
