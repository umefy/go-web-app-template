package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type AppEnv string

const (
	AppEnvDev  AppEnv = "dev"
	AppEnvProd AppEnv = "prod"
)

type appConfig struct {
	Env        AppEnv            `mapstructure:"env"`
	HttpServer *httpServerConfig `mapstructure:"http_server"`
	Logging    *loggingConfig    `mapstructure:"logging"`
	DataBase   *dbConfig         `mapstructure:"database"`
	GrpcServer *grpcServerConfig `mapstructure:"grpc_server"`
}

var _ validation.Validate = (*appConfig)(nil)

func (a *appConfig) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Env, validation.In(AppEnvDev, AppEnvProd)),
		validation.Field(&a.HttpServer, validation.Required),
		validation.Field(&a.Logging, validation.Required),
		validation.Field(&a.DataBase, validation.Required),
		validation.Field(&a.GrpcServer, validation.Required),
	)
}
