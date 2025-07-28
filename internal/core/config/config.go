package config

type Config interface {
	GetAppConfig() AppConfig
	GetEnv() AppEnv
	GetHttpServerConfig() HttpServerConfig
	GetLoggingConfig() LoggingConfig
	GetDBConfig() DbConfig
	GetGrpcServerConfig() GrpcServerConfig
}

type coreConfig struct {
	appConfig AppConfig
}

var _ Config = (*coreConfig)(nil)

func NewAppConfig(appConfig AppConfig) *coreConfig {
	return &coreConfig{appConfig: appConfig}
}

func (c *coreConfig) GetAppConfig() AppConfig {
	return c.appConfig
}

func (c *coreConfig) GetEnv() AppEnv {
	return c.appConfig.Env
}

func (c *coreConfig) GetHttpServerConfig() HttpServerConfig {
	return c.appConfig.HttpServer
}

func (c *coreConfig) GetLoggingConfig() LoggingConfig {
	return c.appConfig.Logging
}

func (c *coreConfig) GetDBConfig() DbConfig {
	return c.appConfig.DataBase
}

func (c *coreConfig) GetGrpcServerConfig() GrpcServerConfig {
	return c.appConfig.GrpcServer
}
