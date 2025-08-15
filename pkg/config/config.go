package config

import (
	"bytes"
	"embed"
	"strings"

	"github.com/spf13/viper"
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type ConfigOption struct {
	ConfigType  string
	ConfigName  string
	ConfigFS    *embed.FS // embed.FS has higher priority than ConfigPaths
	ConfigPaths []string
	EnvPrefix   string
}

func Unmarshal(config validation.Validate, opt ConfigOption) error {
	viper := viper.New()

	viper.SetConfigType(opt.ConfigType)
	viper.SetConfigName(opt.ConfigName)

	viper.SetEnvPrefix(opt.EnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // convert env var that has '.' to '_' when you use viper.Get. eg, viper.GetInt("http_server.port") will look for HTTP_SERVER_PORT env var. In yaml config, it's first convert things to like 'http_server.port', then this will convert it to 'HTTP_SERVER_PORT' to search env var.
	viper.AutomaticEnv()

	// if configFS is not nil, read config from embed.FS
	if opt.ConfigFS != nil {
		bytesData, err := opt.ConfigFS.ReadFile(opt.ConfigName)
		if err != nil {
			return err
		}
		err = viper.ReadConfig(bytes.NewReader(bytesData))
		if err != nil {
			return err
		}
	} else {
		// if configFS is nil, read config from file
		for _, path := range opt.ConfigPaths {
			viper.AddConfigPath(path)
		}
		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}

	err := viper.Unmarshal(config)
	if err != nil {
		return err
	}

	return nil
}
