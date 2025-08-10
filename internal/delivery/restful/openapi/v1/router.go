package v1

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/user"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
	"go.uber.org/fx"
)

type ApiV1RouterParams struct {
	fx.In

	user.UserRouterParams
}

func NewApiV1Router(params ApiV1RouterParams) http.Handler {
	r := router.NewRouter()

	r.Mount("/users", user.NewUserRouter(params.UserRouterParams))

	return r
}
