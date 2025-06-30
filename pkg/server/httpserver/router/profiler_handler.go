package router

import (
	"github.com/go-chi/chi/v5/middleware"
)

var ProfilerHandler = middleware.Profiler()
