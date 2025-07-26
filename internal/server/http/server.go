package http

import (
	"fmt"
	"net/http"

	"github.com/umefy/go-web-app-template/internal/app"
	"github.com/umefy/go-web-app-template/internal/core/config"
	"github.com/umefy/go-web-app-template/internal/delivery/graphql"
	apiV1 "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router/middleware"
)

func New(configOptions config.Options) (*httpserver.Server, error) {
	app, err := app.New(configOptions)
	if err != nil {
		return nil, err
	}

	if !app.Config.GetHttpServerConfig().Enabled {
		return nil, nil
	}

	sever := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", app.Config.GetHttpServerConfig().Port),
		Handler: newHttpHandler(app),
	}

	s := httpserver.New(
		sever,
		app.Logger.GetLogger(),
	)

	return s, nil
}

func newHttpHandler(app *app.App) http.Handler {
	r := router.NewRootRouter(app.Logger.GetLogger())

	r.Use(middleware.Cors(app.Config.GetHttpServerConfig().AllowedOrigins))
	r.Use(middleware.HealthCheck(app.Config.GetHttpServerConfig().HealthCheckEndpoint))

	r.Mount(app.Config.GetHttpServerConfig().ProfilerEndpoint, router.ProfilerHandler)
	r.Mount("/api/v1", apiV1.NewApiV1Router(app))
	r.Mount("/graphql", graphql.NewGraphqlRouter(app))
	return r
}
