package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type HttpServerConfig struct {
	Enabled             bool
	Port                int
	AllowedOrigins      []string `mapstructure:"allowed_origins"`
	HealthCheckEndpoint string   `mapstructure:"health_check_endpoint"`
	ProfilerEndpoint    string   `mapstructure:"profiler_endpoint"`
}

var _ validation.Validate = (*HttpServerConfig)(nil)

func (s *HttpServerConfig) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Enabled, validation.In(true, false).Error("can only be set to true or false")),
		validation.Field(&s.Port, validation.Required),
		validation.Field(&s.AllowedOrigins, validation.Required),
		validation.Field(&s.HealthCheckEndpoint, validation.Required),
		validation.Field(&s.ProfilerEndpoint, validation.Required),
	)
}
