package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/umefy/go-web-app-template/internal/core/config"
	"github.com/umefy/go-web-app-template/internal/delivery/grpc/greeter"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	greeterSvc "github.com/umefy/go-web-app-template/internal/service/greeter"
	"github.com/umefy/go-web-app-template/pkg/server/grpcserver"
	pb "github.com/umefy/go-web-app-template/protogen/grpc/service"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServerParams struct {
	fx.In

	ConfigOptions  config.Options
	Config         config.Config
	Logger         logger.Logger
	GreeterService greeterSvc.Service
}

func registerServices(grpcServer *grpc.Server, params GrpcServerParams) {
	pb.RegisterGreeterServer(grpcServer, greeter.NewHandler(params.Logger, params.GreeterService))
}

func NewServer(params GrpcServerParams) (*grpcserver.GrpcServer, error) {

	if !params.Config.GetGrpcServerConfig().Enabled {
		return nil, nil
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", params.Config.GetGrpcServerConfig().Port))
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryRecoveryInterceptor(params.Logger)),
		grpc.StreamInterceptor(streamRecoveryInterceptor(params.Logger)),
	)

	registerServices(grpcServer, params)
	server := grpcserver.New(listener, grpcServer, params.Logger.GetLogger())
	return server, nil
}

func unaryRecoveryInterceptor(logger logger.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("Recovered from panic in %s: %v", slog.String("method", info.FullMethod), slog.Any("error", r))
				err = status.Errorf(codes.Internal, "Internal server error")
			}
		}()
		return handler(ctx, req)
	}
}

func streamRecoveryInterceptor(logger logger.Logger) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) (err error) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("Recovered from panic in %s: %v", slog.String("method", info.FullMethod), slog.Any("error", r))
				err = status.Errorf(codes.Internal, "Panic in %s: %v", info.FullMethod, r)
			}
		}()
		return handler(srv, stream)
	}
}
