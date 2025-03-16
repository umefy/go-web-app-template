package handler

import (
	"errors"
	"log/slog"
	"net/http"

	bizError "github.com/umefy/go-web-app-template/app/error"
	loggerSrv "github.com/umefy/go-web-app-template/app/logger/service"
	"github.com/umefy/godash/jsonkit"
)

type DefaultHandler struct {
	loggerService loggerSrv.Service
}

var _ Handler = (*DefaultHandler)(nil)

func NewDefaultHandler(loggerService loggerSrv.Service) *DefaultHandler {
	return &DefaultHandler{loggerService: loggerService}
}

func (h *DefaultHandler) HandlerFunc(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			h.HandleError(w, r, err)
		}
	}
}

func (h *DefaultHandler) HandleError(w http.ResponseWriter, r *http.Request, err error) {

	h.loggerService.ErrorContext(r.Context(), "DefaultErrorHandler", slog.String("error", err.Error()))

	var bizErr *bizError.Error
	if errors.As(err, &bizErr) {
		// nolint: errcheck
		jsonkit.JSONResponse(w, bizErr.HTTPCode, map[string]string{"errorCode": bizErr.ErrorCode, "errorMsg": bizErr.ErrorMsg})
		return
	}
	// nolint: errcheck
	jsonkit.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
}
