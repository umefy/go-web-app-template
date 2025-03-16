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

	viper.AutomaticEnv()
	viper.SetEnvPrefix(opt.EnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // convert env var that has '_' to '.' during unmarshal

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
