package config

import (
	"time"

	"github.com/umefy/go-web-app-template/pkg/validation"
)

var DB_LOG_LEVELS = []interface{}{"silent", "error", "warn", "info"}

type DbLoggerConfig struct {
	Writer                 string `mapstructure:"writer"`
	Level                  string `mapstructure:"level"`
	ShowSqlParams          bool   `mapstructure:"show_sql_params"`
	SlowThresholdInSeconds int    `mapstructure:"slow_threshold_in_seconds"`
}

var _ validation.Validate = (*DbLoggerConfig)(nil)

func (c DbLoggerConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Writer, validation.Required, validation.In(LOG_WRITERS...).Error("can only be set to stdout")),
		validation.Field(&c.Level, validation.Required, validation.In(DB_LOG_LEVELS...).Error("can only be set to silent, error, warn, info")),
		validation.Field(&c.ShowSqlParams, validation.In(true, false).Error("can only be set to true or false")),
		validation.Field(&c.SlowThresholdInSeconds, validation.Required, validation.Min(0).Error("must be greater than 0")),
	)
}

type DbConfig struct {
	Url             string         `mapstructure:"url"`
	MaxIdleConns    int            `mapstructure:"max_idle_conns"`
	MaxOpenConns    int            `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration  `mapstructure:"conn_max_lifetime"`
	Logger          DbLoggerConfig `mapstructure:"logger"`
}

var _ validation.Validate = (*DbConfig)(nil)

func (c DbConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Url, validation.Required),
		validation.FieldStruct(&c.Logger),
	)
}

func (c DbConfig) GetLoggerConfig() DbLoggerConfig {
	return c.Logger
}
