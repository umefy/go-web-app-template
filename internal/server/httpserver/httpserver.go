package httpserver

import (
	"fmt"
	"net/http"

	"github.com/umefy/go-web-app-template/internal/app"
	user "github.com/umefy/go-web-app-template/internal/server/httpserver/user"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
)

func New(args app.Arguments) (*httpserver.Server, error) {
	app, err := app.New(args)
	if err != nil {
		return nil, err
	}

	sever := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.ConfigService.GetServerConfig().Port),
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

	r.Mount("/users", user.NewRouter(app))
	return r
}
