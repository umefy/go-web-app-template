package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type AppEnv string

const (
	AppEnvDev  AppEnv = "dev"
	AppEnvProd AppEnv = "prod"
)

type AppConfig struct {
	Env        AppEnv           `mapstructure:"env"`
	HttpServer HttpServerConfig `mapstructure:"http_server"`
	Logging    LoggingConfig    `mapstructure:"logging"`
	DataBase   DbConfig         `mapstructure:"database"`
	GrpcServer GrpcServerConfig `mapstructure:"grpc_server"`
}

var _ validation.Validate = (*AppConfig)(nil)

func (a *AppConfig) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Env, validation.In(AppEnvDev, AppEnvProd)),
		validation.Field(&a.HttpServer, validation.Required),
		validation.Field(&a.Logging, validation.Required),
		validation.Field(&a.DataBase, validation.Required),
		validation.Field(&a.GrpcServer, validation.Required),
	)
}
