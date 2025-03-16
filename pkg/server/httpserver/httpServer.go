package httpserver

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/umefy/godash/logger"
)

type Server struct {
	server *http.Server
	logger *slog.Logger
}

func New(server *http.Server, logger *logger.Logger) *Server {

	loggerHandler := logger.GetHandler()
	loggerHandler.CallerSkip = 3
	slogger := slog.New(&loggerHandler)

	s := &Server{
		server: server,
		logger: slogger,
	}

	return s
}

func (s *Server) Start(shutdownTimeout time.Duration) {
	// Listen for syscall signals for process to interrupt/quit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		s.logger.Info("Starting server", slog.String("address", s.server.Addr))

		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("Server error", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	<-stop
	s.logger.Info("Graceful shutting down the server...", slog.String("timeout", shutdownTimeout.String()))

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("Error shutting down server", slog.String("error", err.Error()))
	} else {
		s.logger.Info("Server gracefully stopped")
	}
}
