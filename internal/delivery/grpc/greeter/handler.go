package greeter

import (
	"context"
	"log/slog"

	greeterSvc "github.com/umefy/go-web-app-template/internal/domain/greeter/service"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	pb "github.com/umefy/go-web-app-template/protogen/grpc/service"
)

type greeterHandler struct {
	logger         logger.Logger
	greeterService greeterSvc.Service
	pb.UnimplementedGreeterServer
}

var _ pb.GreeterServer = (*greeterHandler)(nil)

func NewHandler(logger logger.Logger, greeterService greeterSvc.Service) *greeterHandler {
	return &greeterHandler{logger: logger, greeterService: greeterService}
}

// SayHello implements pb.GreeterServer.
func (g *greeterHandler) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	g.logger.InfoContext(ctx, "SayHello", slog.String("name", req.Name))
	message, err := g.greeterService.SayHello(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.HelloResponse{Message: message}, nil
}
