package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type loggingConfig struct {
	Level string
}

var _ validation.Validate = (*loggingConfig)(nil)

func (l loggingConfig) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Level, validation.In("debug", "info", "warn", "error").Error("can only be set to debug, info, warn, error")),
	)
}
