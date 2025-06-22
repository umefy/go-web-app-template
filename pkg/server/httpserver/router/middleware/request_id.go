package middleware

import (
	"context"
	"fmt"
	"hash/fnv"
	"net/http"

	"github.com/google/uuid"
)

type requestIDKey struct{}

var RequestIDKey = requestIDKey{}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqID := getRequestID(r)
		w.Header().Set("X-Request-ID", reqID) // Attach to response headers

		// Store in context for retrieval
		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getRequestID(r *http.Request) string {
	requestId := r.Header.Get("X-Request-ID")
	if requestId != "" {
		return requestId
	}

	remoteIP := ExtractIP(r)

	u := uuid.New()
	hasher := fnv.New64a()
	hasher.Write([]byte(u.String()))
	hash := hasher.Sum64()
	hashedID := fmt.Sprintf("%x", hash)[:8]

	// Keep the IP visible in the final request ID
	requestId = fmt.Sprintf("%s-%s", remoteIP, hashedID)

	return requestId
}
