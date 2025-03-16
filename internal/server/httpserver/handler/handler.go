package handler

import (
	"net/http"
)

type Handler interface {
	HandlerFunc(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc
}
