package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/restful/handler/middleware"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
	"go.uber.org/fx"
)

type UserRouterParams struct {
	fx.In

	UserService userSvc.Service
	Logger      logger.Logger
	DbQuery     *database.Query
}

func NewUserRouter(params UserRouterParams) http.Handler {

	r := router.NewRouter()

	h := NewHandler(params.UserService, params.Logger)

	r.Get("/", h.Handle(h.GetUsers))
	r.Get("/{id}", h.Handle(h.GetUser))
	r.Post("/", h.Handle(h.ApplyMiddlewares(
		h.CreateUser,
		middleware.Transaction(params.DbQuery, params.Logger),
	)))
	r.Patch("/{id}", h.Handle(h.ApplyMiddlewares(
		h.UpdateUser,
		middleware.Transaction(params.DbQuery, params.Logger),
	)))
	return r
}
