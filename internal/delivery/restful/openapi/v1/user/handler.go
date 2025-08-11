package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/restful/handler"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/handler/middleware"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	userSrv "github.com/umefy/go-web-app-template/internal/service/user"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router"
)

type Handler interface {
	handler.Handler
	handler.Router
	GetUsers(w http.ResponseWriter, r *http.Request) error
	GetUser(w http.ResponseWriter, r *http.Request) error
	CreateUser(w http.ResponseWriter, r *http.Request) error
	UpdateUser(w http.ResponseWriter, r *http.Request) error
}

type userHandler struct {
	*handler.DefaultHandler
	userService userSrv.Service
	logger      logger.Logger
	dbQuery     *database.Query
}

const userHandlerName = "UserHandler"

var _ Handler = (*userHandler)(nil)

func NewHandler(userService userSrv.Service, logger logger.Logger, dbQuery *database.Query) *userHandler {
	return &userHandler{
		DefaultHandler: handler.NewDefaultHandler(
			userHandlerName,
			logger,
		),
		userService: userService,
		logger:      logger,
		dbQuery:     dbQuery,
	}
}

func (h *userHandler) RegisterRoutes(r *router.Mux) {
	r.Route("/users", func(r router.Router) {

		r.Get("/", h.Handle(h.GetUsers))
		r.Get("/{id}", h.Handle(h.GetUser))
		r.Post("/", h.Handle(h.ApplyMiddlewares(
			h.CreateUser,
			middleware.Transaction(h.dbQuery, h.logger),
		)))
		r.Patch("/{id}", h.Handle(h.ApplyMiddlewares(
			h.UpdateUser,
			middleware.Transaction(h.dbQuery, h.logger),
		)))
	})
}

// // Custom error handler
// func (h *userHandler) Handle(handlerFunc handler.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if err := handlerFunc(w, r); err != nil {
// 			h.loggerService.ErrorContext(r.Context(), "UserHandler", slog.String("error", err.Error()))
// 			h.DefaultHandler.HandleError(w, r, err)
// 		}
// 	}
// }
