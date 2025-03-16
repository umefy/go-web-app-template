package service

import (
	"github.com/umefy/go-web-app-template/app/config"
)

type Service interface {
	GetAppConfig() *config.AppConfig
	GetServerConfig() *config.ServerConfig
	GetLoggingConfig() *config.LoggingConfig
	GetDBConfig() *config.DBConfig
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

func (s *service) GetServerConfig() *config.ServerConfig {
	return s.appConfig.Server
}

func (s *service) GetLoggingConfig() *config.LoggingConfig {
	return s.appConfig.Logging
}

func (s *service) GetDBConfig() *config.DBConfig {
	return s.appConfig.DATABASE
}
