package config

import "github.com/umefy/go-web-app-template/pkg/validation"

type grpcServerConfig struct {
	Enabled bool
	Port    int
}

var _ validation.Validate = (*grpcServerConfig)(nil)

func (c *grpcServerConfig) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Enabled, validation.In(true, false).Error("can only be set to true or false")),
		validation.Field(&c.Port, validation.Required),
	)
}
