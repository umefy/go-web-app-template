package handler

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	appError "github.com/umefy/go-web-app-template/internal/domain/error"
	loggerSrv "github.com/umefy/go-web-app-template/internal/domain/logger/service"
	"github.com/umefy/godash/jsonkit"
)

type DefaultHandler struct {
	handlerName   string
	loggerService loggerSrv.Service
}

var _ Handler = (*DefaultHandler)(nil)

func NewDefaultHandler(handlerName string, loggerService loggerSrv.Service) *DefaultHandler {
	if handlerName == "" {
		handlerName = "DefaultHandler"
	}
	return &DefaultHandler{handlerName: handlerName, loggerService: loggerService}
}

func (h *DefaultHandler) ApplyMiddlewares(originalHandler HandlerFunc, middlewares ...Middleware) HandlerFunc {
	handlerFunc := originalHandler
	for _, middleware := range middlewares {
		handlerFunc = middleware(handlerFunc)
	}
	return handlerFunc
}

func (h *DefaultHandler) Handle(handlerFunc HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFunc(w, r); err != nil {
			h.HandleError(w, r, err)
		}
	}
}

func (h *DefaultHandler) HandleError(w http.ResponseWriter, r *http.Request, err error) {
	h.loggerService.ErrorContext(r.Context(), fmt.Sprintf("DefaultHandler(%s) Catch", h.handlerName), slog.String("error", err.Error()))
	var appErr *appError.Error
	if errors.As(err, &appErr) {
		// nolint: errcheck
		jsonkit.JSONResponse(w, appErr.HTTPCode, map[string]string{"errorCode": appErr.ErrorCode, "errorMsg": appErr.ErrorMsg})
		return
	}
	// nolint: errcheck
	jsonkit.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
}
