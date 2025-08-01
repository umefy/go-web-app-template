package config

import (
	"time"

	"github.com/umefy/go-web-app-template/pkg/validation"
)

type DbConfig struct {
	Url             string        `mapstructure:"url"`
	EnableLog       bool          `mapstructure:"enable_log"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

var _ validation.Validate = (*DbConfig)(nil)

func (c *DbConfig) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Url, validation.Required),
		validation.Field(&c.EnableLog, validation.In(true, false).Error("can only be set to true or false")),
	)
}
