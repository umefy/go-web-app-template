package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

func Cors(
	allowedOrigins []string,
) func(next http.Handler) http.Handler {
	return cors.Handler(
		cors.Options{
			AllowedOrigins: allowedOrigins,
		},
	)
}
