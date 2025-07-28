package v1

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/app"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/user"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
)

func NewApiV1Router(app *app.App) http.Handler {
	r := router.NewRouter()

	r.Mount("/users", user.NewUserRouter(app))

	return r
}
