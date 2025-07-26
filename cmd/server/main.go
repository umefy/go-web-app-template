package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/umefy/go-web-app-template/internal/core/config"
	"github.com/umefy/go-web-app-template/internal/server/grpc"
	"github.com/umefy/go-web-app-template/internal/server/http"
	"golang.org/x/sync/errgroup"
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

	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return startHttpServer(args)
	})

	g.Go(func() error {
		return startGrpcServer(args)
	})

	if err := g.Wait(); err != nil {
		log.Fatalln("Failed to start:", err)
	}
}

func startHttpServer(configOptions config.Options) error {
	server, err := http.New(configOptions)

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

func startGrpcServer(configOptions config.Options) error {
	server, err := grpc.New(configOptions)

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
