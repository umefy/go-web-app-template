package service

import (
	"github.com/google/wire"
	pb "github.com/umefy/go-web-app-template/protogen/grpc/service"
)

var WireSet = wire.NewSet(
	NewService,
	wire.Bind(new(pb.GreeterServer), new(*greetService)),
)
