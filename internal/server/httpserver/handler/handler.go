package handler

import (
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

type Handler interface {
	ApplyMiddleware(originalHandler HandlerFunc, middlewares ...Middleware) HandlerFunc
	Handle(handler HandlerFunc) http.HandlerFunc
}

type Middleware func(HandlerFunc) HandlerFunc
