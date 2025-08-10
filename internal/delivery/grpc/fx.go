package grpc

import (
	"github.com/umefy/go-web-app-template/internal/delivery/grpc/greeter"
	pb "github.com/umefy/go-web-app-template/protogen/grpc/service"
	"go.uber.org/fx"
)

var Module = fx.Module("grpcHandler",
	fx.Provide(
		fx.Annotate(
			greeter.NewHandler,
			fx.As(new(pb.GreeterServer)),
		),
	),
)
