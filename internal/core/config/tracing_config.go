package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type TracingConfig struct {
	Enabled        bool
	TracerName     string `mapstructure:"tracer_name"`
	JaegerEndpoint string `mapstructure:"jaeger_endpoint"`
	ServiceName    string `mapstructure:"service_name"`
	ServiceVersion string `mapstructure:"service_version"`
}

var _ validation.Validate = (*TracingConfig)(nil)

func (c TracingConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Enabled, validation.In(true, false).Error("can only be set to true or false")),
		validation.Field(&c.JaegerEndpoint, validation.When(c.Enabled, validation.Required)),
		validation.Field(&c.ServiceName, validation.When(c.Enabled, validation.Required)),
		validation.Field(&c.ServiceVersion, validation.When(c.Enabled, validation.Required)),
		validation.Field(&c.TracerName, validation.When(c.Enabled, validation.Required)),
	)
}
