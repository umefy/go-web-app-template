package user

import (
	"log/slog"
	"net/http"

	loggerSrv "github.com/umefy/go-web-app-template/app/logger/service"
	userSrv "github.com/umefy/go-web-app-template/app/user/service"
	"github.com/umefy/go-web-app-template/internal/server/httpserver/handler"
)

type Handler interface {
	handler.Handler
	GetUsers(w http.ResponseWriter, r *http.Request) error
	GetUser(w http.ResponseWriter, r *http.Request) error
	CreateUser(w http.ResponseWriter, r *http.Request) error
}

type userHandler struct {
	*handler.DefaultHandler
	userService   userSrv.Service
	loggerService loggerSrv.Service
}

var _ Handler = (*userHandler)(nil)

func NewHandler(userService userSrv.Service, loggerService loggerSrv.Service) *userHandler {
	return &userHandler{
		DefaultHandler: handler.NewDefaultHandler(loggerService),
		userService:    userService,
		loggerService:  loggerService,
	}
}

// handleError implements Handler.
func (h *userHandler) HandlerFunc(handle func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handle(w, r); err != nil {
			h.loggerService.ErrorContext(r.Context(), "UserHandler.HandleError", slog.String("error", err.Error()))

			h.HandleError(w, r, err)
		}
	}
}
