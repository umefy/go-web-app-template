package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/umefy/go-web-app-template/internal/core/config"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router/middleware"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

type ServerParams struct {
	fx.In

	Ctx            context.Context // context used for initial different infra
	Config         config.Config
	Logger         logger.Logger
	TracerProvider trace.TracerProvider
	GraphqlRouter  http.Handler `name:"graphqlRouter"`
	ApiV1Router    http.Handler `name:"apiV1Router"`
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

	appConfig := params.Config.GetAppConfig()
	r.Use(middleware.Cors(params.Config.GetHttpServerConfig().AllowedOrigins))
	r.Use(middleware.HealthCheck(params.Config.GetHttpServerConfig().HealthCheckEndpoint, string(appConfig.Env), appConfig.Version, params.Logger.GetLogger()))
	r.Use(middleware.OTelTracing(params.Config.GetHttpServerConfig().ServerName, params.TracerProvider))

	r.Mount(params.Config.GetHttpServerConfig().ProfilerEndpoint, router.ProfilerHandler)
	r.Mount("/api/v1", params.ApiV1Router)
	r.Mount("/graphql", params.GraphqlRouter)
	return r
}
