package middleware

import (
	"net/http"
	"time"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func Timeout(t time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if IsWebSocketUpgrade(r) {
				next.ServeHTTP(w, r)
				return
			}

			chiMiddleware.Timeout(t)(next).ServeHTTP(w, r)
		})
	}
}
