package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/umefy/go-web-app-template/internal/core/config"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	"github.com/umefy/go-web-app-template/internal/infrastructure/server/grpc"
	"github.com/umefy/go-web-app-template/internal/infrastructure/server/http"
	"github.com/umefy/go-web-app-template/internal/infrastructure/tracing"
	"github.com/umefy/go-web-app-template/internal/service"
	"github.com/umefy/go-web-app-template/pkg/server/grpcserver"
	"github.com/umefy/go-web-app-template/pkg/server/httpserver"
	"go.uber.org/fx"
)

func main() {

	var env string
	var name string
	var configPath string
	flag.StringVar(&env, "env", "dev", "active environment. Available options: dev, test, prod.")
	flag.StringVar(&name, "name", "webapp", "app name")
	flag.StringVar(&configPath, "config", "", "config file path. If set, will ignore env option")
	flag.Parse()

	args := config.Options{
		Env:        env,
		ConfigPath: configPath,
	}

	app := fx.New(
		fx.Supply(args),
		fx.Provide(func() context.Context {
			return context.Background()
		}),
		config.Module,
		database.Module,
		logger.Module,
		tracing.Module,
		http.Module,
		grpc.Module,
		service.Module,
		fx.Invoke(start),
	)

	app.Run()
}

func start(ctx context.Context, lc fx.Lifecycle, httpServer *httpserver.Server, grpcServer *grpcserver.GrpcServer, cfg config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start servers in background goroutines without waiting
			go func() {
				if err := startHttpServer(httpServer, cfg); err != nil {
					log.Printf("HTTP server error: %v", err)
				}
			}()

			go func() {
				if err := startGrpcServer(grpcServer, cfg); err != nil {
					log.Printf("gRPC server error: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Server stopped...")
			if err := stopHttpServer(ctx, httpServer, cfg); err != nil {
				log.Printf("HTTP server stop error: %v", err)
				return err
			}

			if err := stopGrpcServer(ctx, grpcServer, cfg); err != nil {
				log.Printf("gRPC server stop error: %v", err)
				return err
			}

			return nil
		},
	})
}

func startHttpServer(httpServer *httpserver.Server, cfg config.Config) error {
	httpCfg := cfg.GetHttpServerConfig()
	if !httpCfg.Enabled {
		log.Println("Http server is not enabled")
		return nil
	}

	return httpServer.Start()
}

func stopHttpServer(ctx context.Context, httpServer *httpserver.Server, cfg config.Config) error {
	httpCfg := cfg.GetHttpServerConfig()
	if !httpCfg.Enabled {
		log.Println("Http server is not enabled")
		return nil
	}

	return httpServer.Shutdown(ctx, time.Duration(httpCfg.ShutdownTimeoutInSeconds)*time.Second)
}

func startGrpcServer(server *grpcserver.GrpcServer, cfg config.Config) error {

	grpcCfg := cfg.GetGrpcServerConfig()
	if !grpcCfg.Enabled {
		log.Println("Grpc server is not enabled")
		return nil
	}

	return server.Start()
}

func stopGrpcServer(ctx context.Context, server *grpcserver.GrpcServer, cfg config.Config) error {
	grpcCfg := cfg.GetGrpcServerConfig()
	if !grpcCfg.Enabled {
		log.Println("Grpc server is not enabled")
		return nil
	}

	return server.Shutdown(ctx, time.Duration(grpcCfg.ShutdownTimeoutInSeconds)*time.Second)
}
