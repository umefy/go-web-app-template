package grpcserver

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"os/signal"
	"syscall"
	"time"

	"github.com/umefy/godash/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	logger   *slog.Logger
	server   *grpc.Server
	listener net.Listener
}

// New creates a new gRPC server.
// It need pass the original grpc.Server and logger.Logger.
// The original grpc.Server can be created by User requirement itself.
// Our wrapper will help to start the server and graceful shutdown.
func New(listener net.Listener, server *grpc.Server, logger *logger.Logger) *GrpcServer {
	loggerHandler := logger.GetHandler()
	loggerHandler.CallerSkip = 3
	slogger := slog.New(&loggerHandler)

	s := &GrpcServer{
		server:   server,
		logger:   slogger,
		listener: listener,
	}

	return s
}

func (s *GrpcServer) Start() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		s.logger.Info("Starting GRPC server", slog.String("address", s.listener.Addr().String()))

		if err := s.server.Serve(s.listener); err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		s.logger.Info("GRPC server shutting down...")
		s.server.GracefulStop()
		return nil
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		s.logger.Error("GRPC server error", slog.String("error", err.Error()))
	}

	s.logger.Info("GRPC server shutdown complete")
}

func (s *GrpcServer) Shutdown(ctx context.Context, timeout time.Duration) error {
	s.logger.Info("Graceful shutting down the GRPC server...", slog.String("timeout", timeout.String()))

	_, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	s.server.GracefulStop()
	return nil
}
