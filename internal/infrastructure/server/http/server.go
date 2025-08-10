package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/umefy/go-web-app-template/internal/core/config"
	"github.com/umefy/go-web-app-template/internal/delivery/graphql"
	apiV1 "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1"
	userRouter "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/user"
	orderRepo "github.com/umefy/go-web-app-template/internal/domain/order/repo"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repo"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	greeterSvc "github.com/umefy/go-web-app-template/internal/service/greeter"
	orderSvc "github.com/umefy/go-web-app-template/internal/service/order"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router/middleware"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

type ServerParams struct {
	fx.In

	Ctx             context.Context // context used for initial different infra
	Arguments       config.Options
	Config          config.Config
	Logger          logger.Logger
	UserService     userSvc.Service
	UserRepository  userRepo.Repository
	OrderService    orderSvc.Service
	OrderRepository orderRepo.Repository
	GreeterService  greeterSvc.Service
	DbQuery         *database.Query
	Tracer          trace.Tracer
	ConfigOptions   config.Options
}

func NewServer(params ServerParams) (*httpserver.Server, error) {

	if !params.Config.GetHttpServerConfig().Enabled {
		return nil, nil
	}

	sever := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", params.Config.GetHttpServerConfig().Port),
		Handler: newHttpHandler(params),
	}

	s := httpserver.New(
		sever,
		params.Logger.GetLogger(),
	)

	return s, nil
}

func newHttpHandler(params ServerParams) http.Handler {
	r := router.NewRootRouter(params.Logger.GetLogger())

	r.Use(middleware.Cors(params.Config.GetHttpServerConfig().AllowedOrigins))
	r.Use(middleware.HealthCheck(params.Config.GetHttpServerConfig().HealthCheckEndpoint))
	r.Use(middleware.OTelTracing(params.Config.GetTracingConfig().TracerName, params.Tracer))

	r.Mount(params.Config.GetHttpServerConfig().ProfilerEndpoint, router.ProfilerHandler)
	r.Mount("/api/v1", apiV1.NewApiV1Router(apiV1.ApiV1RouterParams{
		UserRouterParams: userRouter.UserRouterParams{
			UserService: params.UserService,
			Logger:      params.Logger,
			DbQuery:     params.DbQuery,
		},
	}))
	r.Mount("/graphql", graphql.NewGraphqlRouter(graphql.GraphqlRouterParams{
		UserService:  params.UserService,
		Logger:       params.Logger,
		Config:       params.Config,
		DbQuery:      params.DbQuery,
		OrderService: params.OrderService,
	}))
	return r
}
