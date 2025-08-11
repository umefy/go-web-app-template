package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type HttpServerConfig struct {
	Enabled             bool
	Port                int
	ServerName          string   `mapstructure:"server_name"`
	AllowedOrigins      []string `mapstructure:"allowed_origins"`
	HealthCheckEndpoint string   `mapstructure:"health_check_endpoint"`
	ProfilerEndpoint    string   `mapstructure:"profiler_endpoint"`
}

var _ validation.Validate = (*HttpServerConfig)(nil)

func (s HttpServerConfig) Validate() error {

	return validation.ValidateStruct(&s,
		validation.Field(&s.Enabled, validation.In(true, false).Error("can only be set to true or false")),
		validation.Field(&s.ServerName, validation.When(s.Enabled, validation.Required)),
		validation.Field(&s.Port, validation.When(s.Enabled, validation.Required)),
		validation.Field(&s.AllowedOrigins, validation.When(s.Enabled, validation.Required)),
		validation.Field(&s.HealthCheckEndpoint, validation.When(s.Enabled, validation.Required)),
		validation.Field(&s.ProfilerEndpoint, validation.When(s.Enabled, validation.Required)),
	)
}
