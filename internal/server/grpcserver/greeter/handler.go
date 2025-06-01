package greeter

import (
	"context"
	"log/slog"

	greeterSvc "github.com/umefy/go-web-app-template/app/greeter/service"
	loggerSrv "github.com/umefy/go-web-app-template/app/logger/service"
	pb "github.com/umefy/go-web-app-template/protogen/grpc/service"
)

type greeterHandler struct {
	loggerService  loggerSrv.Service
	greeterService greeterSvc.Service
	pb.UnimplementedGreeterServer
}

var _ pb.GreeterServer = (*greeterHandler)(nil)

func NewHandler(loggerService loggerSrv.Service, greeterService greeterSvc.Service) *greeterHandler {
	return &greeterHandler{loggerService: loggerService, greeterService: greeterService}
}

// SayHello implements pb.GreeterServer.
func (g *greeterHandler) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	g.loggerService.InfoContext(ctx, "SayHello", slog.String("name", req.Name))
	message, err := g.greeterService.SayHello(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.HelloResponse{Message: message}, nil
}
