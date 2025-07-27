package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/errutil"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	"github.com/umefy/godash/jsonkit"
)

type DefaultHandler struct {
	handlerName string
	logger      logger.Logger
}

var _ Handler = (*DefaultHandler)(nil)

func NewDefaultHandler(handlerName string, logger logger.Logger) *DefaultHandler {
	if handlerName == "" {
		handlerName = "DefaultHandler"
	}
	return &DefaultHandler{handlerName: handlerName, logger: logger}
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
	h.logger.ErrorContext(r.Context(), fmt.Sprintf("DefaultHandler(%s) Catch", h.handlerName), slog.String("error", err.Error()))

	statusCode, errMap := errutil.FormatError(err)
	// nolint: errcheck
	jsonkit.JSONResponse(w, statusCode, errMap)
}
