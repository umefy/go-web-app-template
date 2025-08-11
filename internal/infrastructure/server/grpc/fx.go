package grpc

import (
	grpcHandler "github.com/umefy/go-web-app-template/internal/delivery/grpc"
	"go.uber.org/fx"
)

var Module = fx.Module("grpcServer",
	grpcHandler.Module,
	fx.Provide(NewServer),
)
