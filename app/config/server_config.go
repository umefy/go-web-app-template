package config

import (
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type ServerConfig struct {
	Port int
}

var _ validation.Validate = (*ServerConfig)(nil)

func (s *ServerConfig) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Port, validation.Required),
	)
}
