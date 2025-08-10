package v1

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/restful/handler"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
	"go.uber.org/fx"
)

type ApiV1RouterParams struct {
	fx.In
	Routers []handler.Router `group:"apiV1Routers"`
}

func NewApiV1Router(params ApiV1RouterParams) http.Handler {
	r := router.NewRouter()

	for _, router := range params.Routers {
		router.RegisterRoutes(r)
	}

	return r
}
