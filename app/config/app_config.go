package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type AppConfig struct {
	Server   *ServerConfig
	Logging  *LoggingConfig
	DATABASE *DBConfig
}

var _ validation.Validate = (*AppConfig)(nil)

func (a *AppConfig) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Server, validation.Required),
		validation.Field(&a.Logging, validation.Required),
		validation.Field(&a.DATABASE, validation.Required),
	)
}
