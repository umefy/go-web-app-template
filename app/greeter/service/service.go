package service

import (
	"context"
	"log/slog"

	loggerSrv "github.com/umefy/go-web-app-template/app/logger/service"
	pb "github.com/umefy/go-web-app-template/protogen/grpc/service"
)

type greetService struct {
	loggerService loggerSrv.Service
	pb.UnimplementedGreeterServer
}

var _ pb.GreeterServer = (*greetService)(nil)

func NewService(loggerService loggerSrv.Service) *greetService {
	return &greetService{loggerService: loggerService}
}

func (s *greetService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	s.loggerService.Info("SayHello", slog.String("name", req.Name))
	return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}
