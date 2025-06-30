package http

import (
	"fmt"
	"net/http"

	"github.com/umefy/go-web-app-template/internal/app"
	"github.com/umefy/go-web-app-template/internal/infrastructure/config"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/graphql"
	apiV1 "github.com/umefy/go-web-app-template/internal/infrastructure/http/openapi/v1"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router/middleware"
)

func New(configOptions config.Options) (*httpserver.Server, error) {
	app, err := app.New(configOptions)
	if err != nil {
		return nil, err
	}

	if !app.ConfigService.GetHttpServerConfig().Enabled {
		return nil, nil
	}

	sever := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", app.ConfigService.GetHttpServerConfig().Port),
		Handler: newHttpHandler(app),
	}

	s := httpserver.New(
		sever,
		app.Logger,
	)

	return s, nil
}

func newHttpHandler(app *app.App) http.Handler {
	r := router.NewRootRouter(app.Logger)

	r.Use(middleware.Cors(app.ConfigService.GetHttpServerConfig().AllowedOrigins))
	r.Use(middleware.HealthCheck(app.ConfigService.GetHttpServerConfig().HealthCheckEndpoint))

	r.Mount(app.ConfigService.GetHttpServerConfig().ProfilerEndpoint, router.ProfilerHandler)
	r.Mount("/api/v1", apiV1.NewApiV1Router(app))
	r.Mount("/graphql", graphql.NewGraphqlRouter(app))
	return r
}
