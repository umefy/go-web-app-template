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
	greeterSvc "github.com/umefy/go-web-app-template/internal/service/greeter"
	orderSvc "github.com/umefy/go-web-app-template/internal/service/order"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
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
		userSvc.Module,
		orderSvc.Module,
		greeterSvc.Module,
		fx.Invoke(start),
	)

	app.Run()

}

func start(ctx context.Context, lc fx.Lifecycle, httpParams http.ServerParams, grpcParams grpc.GrpcServerParams) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start servers in background goroutines without waiting
			go func() {
				if err := startHttpServer(httpParams); err != nil {
					log.Printf("HTTP server error: %v", err)
				}
			}()

			go func() {
				if err := startGrpcServer(grpcParams); err != nil {
					log.Printf("gRPC server error: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Server stopped...")
			return nil
		},
	})
}

func startHttpServer(params http.ServerParams) error {
	server, err := http.NewServer(params)

	if err != nil {
		return err
	}

	if server == nil {
		log.Println("Http server is not enabled")
		return nil
	}

	server.Start(10 * time.Second)

	return nil
}

func startGrpcServer(params grpc.GrpcServerParams) error {
	server, err := grpc.NewServer(params)

	if err != nil {
		return err
	}

	if server == nil {
		log.Println("Grpc server is not enabled")
		return nil
	}

	server.Start()

	return nil
}
