package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/umefy/godash/jsonkit"
	"github.com/umefy/godash/logger"
)

func HealthCheck(endpoint string, env string, version string, logger *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if (r.Method == http.MethodGet || r.Method == http.MethodHead) && strings.EqualFold(r.URL.Path, endpoint) {
				resp := map[string]string{
					"status":  "ok",
					"env":     env,
					"version": version,
				}
				err := jsonkit.JSONResponse(w, http.StatusOK, resp)
				if err != nil {
					logger.ErrorContext(r.Context(), "Health check failed", slog.String("error", err.Error()))
				}
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
