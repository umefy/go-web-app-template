package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type GrpcServerConfig struct {
	Enabled bool
	Port    int
}

var _ validation.Validate = (*GrpcServerConfig)(nil)

func (c GrpcServerConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Enabled, validation.In(true, false).Error("can only be set to true or false")),
		validation.Field(&c.Port, validation.When(c.Enabled, validation.Required)),
	)
}
