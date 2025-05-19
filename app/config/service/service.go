package service

import (
	"github.com/umefy/go-web-app-template/app/config"
)

type Service interface {
	GetAppConfig() *config.AppConfig
	GetHttpServerConfig() *config.HttpServerConfig
	GetLoggingConfig() *config.LoggingConfig
	GetDBConfig() *config.DBConfig
	GetGrpcServerConfig() *config.GrpcServerConfig
}

type service struct {
	appConfig *config.AppConfig
}

var _ Service = (*service)(nil)

func NewService(config *config.AppConfig) *service {
	return &service{
		appConfig: config,
	}
}

func (s *service) GetAppConfig() *config.AppConfig {
	return s.appConfig
}

func (s *service) GetHttpServerConfig() *config.HttpServerConfig {
	return s.appConfig.HttpServer
}

func (s *service) GetLoggingConfig() *config.LoggingConfig {
	return s.appConfig.Logging
}

func (s *service) GetDBConfig() *config.DBConfig {
	return s.appConfig.DATABASE
}

func (s *service) GetGrpcServerConfig() *config.GrpcServerConfig {
	return s.appConfig.GrpcServer
}
