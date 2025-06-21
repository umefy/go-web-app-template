package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/app"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/middleware"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
)

func NewRouter(app *app.App) http.Handler {

	r := router.NewRouter()

	h := NewHandler(app.UserService, app.LoggerService)

	r.Get("/", h.Handle(h.GetUsers))
	r.Get("/{id}", h.Handle(h.GetUser))
	r.Post("/", h.Handle(h.CreateUser))
	r.Patch("/{id}", h.Handle(h.ApplyMiddlewares(
		h.UpdateUser,
		middleware.Transaction(app.DbQuery, app.LoggerService),
	)))
	return r
}
