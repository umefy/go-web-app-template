package config

import (
	"fmt"
	"path/filepath"
	"strings"

	yamlConfigs "github.com/umefy/go-web-app-template/configs"
	"github.com/umefy/go-web-app-template/pkg/config"
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type Options struct {
	Env        string
	ConfigPath string
}

const configDirRelativePath = "../../../configs"

var _ validation.Validate = (*Options)(nil)

func (o *Options) Validate() error {
	return validation.ValidateStruct(o,
		validation.Field(&o.Env, validation.In("dev", "test", "prod").Error("can only be set to dev, test or prod")),
	)
}

func NewConfig(args Options) (Config, error) {
	err := args.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalidate config options: %w", err)
	}

	env := args.Env
	configPath := args.ConfigPath

	var opt config.ConfigOption
	if configPath != "" {
		opt = getConfigOptByPath(configPath)
	} else {
		opt = getConfigOptByEnv(env)
	}

	var appConfig AppConfig
	err = config.Unmarshal(&appConfig, opt)

	if err != nil {
		return nil, fmt.Errorf("unmarshal config error: %w", err)
	}

	err = appConfig.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate config error: %w", err)
	}
	return NewAppConfig(appConfig), nil
}

func getConfigOptByEnv(env string) config.ConfigOption {
	configFS := yamlConfigs.FS

	return config.ConfigOption{
		ConfigType: "yaml",
		ConfigName: fmt.Sprintf("app-%s.yaml", env),
		ConfigFS:   &configFS,
		EnvPrefix:  "",
	}
}

func getConfigOptByPath(configPath string) config.ConfigOption {
	configDir := filepath.Dir(configPath)
	configExt := filepath.Ext(configPath)
	configExt = strings.TrimPrefix(configExt, ".")
	configName := filepath.Base(configPath)

	return config.ConfigOption{
		ConfigType:  configExt,
		ConfigName:  configName,
		ConfigPaths: []string{configDir},
		EnvPrefix:   "",
	}
}
