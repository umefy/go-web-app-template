package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

var LOG_LEVELS = []interface{}{"debug", "info", "warn", "error"}

var LOG_WRITERS = []interface{}{"stdout"}

type LoggingConfig struct {
	Level     string
	Writer    string `mapstructure:"writer"`
	UseJson   bool   `mapstructure:"use_json"`
	AddSource bool   `mapstructure:"add_source"` // Add logging source file and line number
	SourceKey string `mapstructure:"source_key"` // The source key field name in the log
}

var _ validation.Validate = (*LoggingConfig)(nil)

func (l LoggingConfig) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Level, validation.Required, validation.In(LOG_LEVELS...).Error("can only be set to debug, info, warn, error")),
		validation.Field(&l.Writer, validation.In(LOG_WRITERS...).Error("can only be set to stdout")),
		validation.Field(&l.UseJson, validation.In(true, false).Error("can only be set to true or false")),
		validation.Field(&l.AddSource, validation.In(true, false).Error("can only be set to true or false")),
		validation.Field(&l.SourceKey, validation.When(l.AddSource, validation.Required)),
	)
}
