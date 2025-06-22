package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type LoggingConfig struct {
	Level string
}

var _ validation.Validate = (*LoggingConfig)(nil)

func (l LoggingConfig) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Level, validation.In("debug", "info", "warn", "error").Error("can only be set to debug, info, warn, error")),
	)
}
