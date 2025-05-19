package config

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type ConfigOption struct {
	ConfigType  string
	ConfigName  string
	ConfigPaths []string
	EnvPrefix   string
}

func Unmarshal(opt ConfigOption, config validation.Validate) error {
	viper := viper.New()

	viper.SetConfigType(opt.ConfigType)
	viper.SetConfigName(opt.ConfigName)
	for _, path := range opt.ConfigPaths {
		viper.AddConfigPath(path)
	}

	viper.SetEnvPrefix(opt.EnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // convert env var that has '.' to '_' when you use viper.Get. eg, viper.GetInt("http_server.port") will look for HTTP_SERVER_PORT env var. In yaml config, it's first convert things to like 'http_server.port', then this will convert it to 'HTTP_SERVER_PORT' to search env var.
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return err
	}

	return nil
}
