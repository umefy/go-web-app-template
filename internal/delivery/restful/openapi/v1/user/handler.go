package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/restful/handler"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	userSrv "github.com/umefy/go-web-app-template/internal/service/user"
)

type Handler interface {
	handler.Handler
	GetUsers(w http.ResponseWriter, r *http.Request) error
	GetUser(w http.ResponseWriter, r *http.Request) error
	CreateUser(w http.ResponseWriter, r *http.Request) error
	UpdateUser(w http.ResponseWriter, r *http.Request) error
}

type userHandler struct {
	*handler.DefaultHandler
	userService userSrv.Service
	logger      logger.Logger
}

const userHandlerName = "UserHandler"

var _ Handler = (*userHandler)(nil)

func NewHandler(userService userSrv.Service, logger logger.Logger) *userHandler {
	return &userHandler{
		DefaultHandler: handler.NewDefaultHandler(
			userHandlerName,
			logger,
		),
		userService: userService,
		logger:      logger,
	}
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
