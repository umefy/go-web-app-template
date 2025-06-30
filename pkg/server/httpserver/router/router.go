package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver/router/middleware"
	"github.com/umefy/godash/logger"
)

var (
	allowedContentTypes = [...]string{"application/json"}
)

func NewRootRouter(logger *logger.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger(logger))
	r.Use(middleware.Recover(logger))

	r.Use(chiMiddleware.AllowContentType(allowedContentTypes[:]...))
	r.Use(chiMiddleware.Timeout(time.Second * 60))
	r.Use(httprate.LimitAll(600, time.Minute))
	r.Use(httprate.LimitByIP(100, time.Minute))

	r.Use(chiMiddleware.Heartbeat("/health-check"))
	r.Mount("/debug", chiMiddleware.Profiler())
	return r
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	return r
}
