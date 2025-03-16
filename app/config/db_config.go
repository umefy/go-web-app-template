package config

import (
	"time"

	"github.com/umefy/go-web-app-template/pkg/validation"
)

type DBConfig struct {
	Url             string        `mapstructure:"url"`
	EnableLog       bool          `mapstructure:"enable_log"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

var _ validation.Validate = (*DBConfig)(nil)

func (c *DBConfig) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Url, validation.Required),
	)
}
