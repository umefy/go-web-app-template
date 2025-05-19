package httpserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/umefy/godash/logger"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	server *http.Server
	logger *slog.Logger
}

// New creates a new http server.
// It need pass the original http.Server and logger.Logger.
// The original http.Server can be created by User requirement itself.
// Our wrapper will help to start the server and graceful shutdown.
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
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		s.logger.Info("Starting HTTP server", slog.String("address", s.server.Addr))

		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("HTTP server error", slog.String("error", err.Error()))
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		s.logger.Info("Graceful shutting down the HTTP server...", slog.String("timeout", shutdownTimeout.String()))

		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
			s.logger.Error("Error shutting down HTTP server", slog.String("error", err.Error()))
			return err
		}
		s.logger.Info("HTTP server gracefully stopped")
		return nil
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		s.logger.Error("Server error", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
