package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/app"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
)

func NewRouter(app *app.App) http.Handler {

	r := router.NewRouter()

	h := NewHandler(app.UserService, app.LoggerService)

	r.Get("/", h.HandlerFunc(h.GetUsers))
	r.Get("/{id}", h.HandlerFunc(h.GetUser))
	r.Post("/", h.HandlerFunc(h.CreateUser))
	r.Patch("/{id}", h.HandlerFunc(h.UpdateUser))
	return r
}
