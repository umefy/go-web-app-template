package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type AppConfig struct {
	HttpServer *HttpServerConfig `mapstructure:"httpServer"`
	Logging    *LoggingConfig    `mapstructure:"logging"`
	DATABASE   *DBConfig         `mapstructure:"database"`
	GrpcServer *GrpcServerConfig `mapstructure:"grpcServer"`
}

var _ validation.Validate = (*AppConfig)(nil)

func (a *AppConfig) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.HttpServer, validation.Required),
		validation.Field(&a.Logging, validation.Required),
		validation.Field(&a.DATABASE, validation.Required),
		validation.Field(&a.GrpcServer, validation.Required),
	)
}
