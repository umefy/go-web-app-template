package v1

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/app"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/middleware"
	userHandler "github.com/umefy/go-web-app-template/internal/infrastructure/http/openapi/v1/handler/user"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
)

func NewApiV1Router(app *app.App) http.Handler {
	r := router.NewRouter()

	r.Mount("/users", newUserRouter(app))

	return r
}

func newUserRouter(app *app.App) http.Handler {

	r := router.NewRouter()

	h := userHandler.NewHandler(app.UserService, app.LoggerService)

	r.Get("/", h.Handle(h.GetUsers))
	r.Get("/{id}", h.Handle(h.GetUser))
	r.Post("/", h.Handle(h.ApplyMiddlewares(
		h.CreateUser,
		middleware.Transaction(app.DbQuery, app.LoggerService),
	)))
	r.Patch("/{id}", h.Handle(h.ApplyMiddlewares(
		h.UpdateUser,
		middleware.Transaction(app.DbQuery, app.LoggerService),
	)))
	return r
}
