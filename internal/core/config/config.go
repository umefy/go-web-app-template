package config

type Config interface {
	GetAppConfig() *appConfig
	GetEnv() AppEnv
	GetHttpServerConfig() *httpServerConfig
	GetLoggingConfig() *loggingConfig
	GetDBConfig() *dbConfig
	GetGrpcServerConfig() *grpcServerConfig
}

type coreConfig struct {
	appConfig *appConfig
}

var _ Config = (*coreConfig)(nil)

func NewAppConfig(appConfig *appConfig) *coreConfig {
	return &coreConfig{appConfig: appConfig}
}

func (c *coreConfig) GetAppConfig() *appConfig {
	return c.appConfig
}

func (c *coreConfig) GetEnv() AppEnv {
	return c.appConfig.Env
}

func (c *coreConfig) GetHttpServerConfig() *httpServerConfig {
	return c.appConfig.HttpServer
}

func (c *coreConfig) GetLoggingConfig() *loggingConfig {
	return c.appConfig.Logging
}

func (c *coreConfig) GetDBConfig() *dbConfig {
	return c.appConfig.DataBase
}

func (c *coreConfig) GetGrpcServerConfig() *grpcServerConfig {
	return c.appConfig.GrpcServer
}
