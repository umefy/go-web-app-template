package handler

import "github.com/umefy/go-web-app-template/pkg/server/httpserver/router"

type Router interface {
	RegisterRoutes(r *router.Mux)
}
