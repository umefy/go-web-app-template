package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/app"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/handler/middleware"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
)

func NewUserRouter(app *app.App) http.Handler {

	r := router.NewRouter()

	h := NewHandler(app.UserService, app.Logger)

	r.Get("/", h.Handle(h.GetUsers))
	r.Get("/{id}", h.Handle(h.GetUser))
	r.Post("/", h.Handle(h.ApplyMiddlewares(
		h.CreateUser,
		middleware.Transaction(app.DbQuery, app.Logger),
	)))
	r.Patch("/{id}", h.Handle(h.ApplyMiddlewares(
		h.UpdateUser,
		middleware.Transaction(app.DbQuery, app.Logger),
	)))
	return r
}
