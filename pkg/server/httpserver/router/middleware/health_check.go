package middleware

import (
	"net/http"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func HealthCheck(endpoint string) func(next http.Handler) http.Handler {
	return chiMiddleware.Heartbeat(endpoint)
}
