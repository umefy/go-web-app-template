package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/umefy/godash/logger"
)

var (
	allowedContentTypes = [...]string{"application/json"}
)

func NewRootRouter(logger *logger.Logger) *chi.Mux {
	r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	r.Use(RequestIDMiddleware)
	r.Use(LoggerMiddleware(logger))

	r.Use(middleware.Recoverer)

	r.Use(middleware.AllowContentType(allowedContentTypes[:]...))
	r.Use(middleware.Timeout(time.Second * 60))
	r.Use(httprate.LimitAll(600, time.Minute))
	r.Use(httprate.LimitByIP(100, time.Minute))

	r.Use(middleware.Heartbeat("/health-check"))
	r.Mount("/debug", middleware.Profiler())
	return r
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	return r
}
