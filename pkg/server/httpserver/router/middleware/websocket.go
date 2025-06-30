package middleware

import (
	"net/http"
	"strings"
)

// IsWebSocketUpgrade checks if the request is attempting to upgrade to WebSocket
func IsWebSocketUpgrade(r *http.Request) bool {
	return strings.ToLower(r.Header.Get("Upgrade")) == "websocket" &&
		strings.ToLower(r.Header.Get("Connection")) == "upgrade"
}
