package errutil

import (
	"errors"
	"net/http"

	domainError "github.com/umefy/go-web-app-template/internal/domain/error"
	"github.com/umefy/go-web-app-template/pkg/validation"
)

func FormatError(err error) (int, map[string]any) {
	var domainErr *domainError.Error
	var validateErr *validation.ValidateStructError

	if errors.As(err, &validateErr) {
		return http.StatusBadRequest, map[string]any{"error": map[string]any{
			"code":    "VALIDATION_ERROR",
			"message": validateErr.Error(),
			"details": validateErr.Errors,
		}}
	}

	if errors.As(err, &domainErr) {
		return domainErr.HTTPCode, map[string]any{"error": map[string]any{
			"code":    domainErr.Code,
			"message": domainErr.Message,
		}}
	}

	return http.StatusInternalServerError, map[string]any{
		"error": map[string]any{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
		},
	}
}
