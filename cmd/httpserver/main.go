package main

import (
	"flag"
	"log"
	"time"

	"github.com/umefy/go-web-app-template/internal/app"
	"github.com/umefy/go-web-app-template/internal/server/httpserver"
)

func main() {

	var env string
	var name string
	var configPath string
	flag.StringVar(&env, "env", "dev", "active environment. Available options: dev, test, prod.")
	flag.StringVar(&name, "name", "webapp", "app name")
	flag.StringVar(&configPath, "config", "", "config file path. If set, will ignore env option")
	flag.Parse()

	args := app.Arguments{
		Env:        env,
		ConfigPath: configPath,
	}
	server, err := httpserver.New(args)
	if err != nil {
		log.Fatalln("Failed to create server:", err)
	}
	server.Start(10 * time.Second)
}
