package user

import (
	"net/http"

	loggerSrv "github.com/umefy/go-web-app-template/internal/domain/logger/service"
	userSrv "github.com/umefy/go-web-app-template/internal/domain/user/service"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/handler"
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
	userService   userSrv.Service
	loggerService loggerSrv.Service
}

const userHandlerName = "UserHandler"

var _ Handler = (*userHandler)(nil)

func NewHandler(userService userSrv.Service, loggerService loggerSrv.Service) *userHandler {
	return &userHandler{
		DefaultHandler: handler.NewDefaultHandler(
			userHandlerName,
			loggerService,
		),
		userService:   userService,
		loggerService: loggerService,
	}
}

// // Custom error handler
// func (h *userHandler) Handle(handlerFunc handler.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if err := handlerFunc(w, r); err != nil {
// 			h.loggerService.ErrorContext(r.Context(), "UserHandler Catch", slog.String("error", err.Error()))

// 			h.DefaultHandler.HandleError(w, r, err)
// 		}
// 	}
// }
