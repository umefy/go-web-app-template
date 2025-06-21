package config

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	appConfig "github.com/umefy/go-web-app-template/internal/domain/config"
	"github.com/umefy/go-web-app-template/pkg/config"
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type Options struct {
	Env        string
	ConfigPath string
}

var _ validation.Validate = (*Options)(nil)

func (o *Options) Validate() error {
	return validation.ValidateStruct(o,
		validation.Field(&o.Env, validation.In("dev", "test", "prod").Error("can only be set to dev, test or prod")),
	)
}

func LoadConfig(args Options) (*appConfig.AppConfig, error) {
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

	var appConfig appConfig.AppConfig
	err = config.Unmarshal(opt, &appConfig)

	if err != nil {
		return nil, fmt.Errorf("unmarshal config error: %w", err)
	}

	err = appConfig.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate config error: %w", err)
	}
	return &appConfig, nil
}

func getConfigOptByEnv(env string) config.ConfigOption {
	_, path, _, _ := runtime.Caller(1)
	configDir := filepath.Join(filepath.Dir(path), "../../configs")

	return config.ConfigOption{
		ConfigType:  "yaml",
		ConfigName:  fmt.Sprintf("app-%s", env),
		ConfigPaths: []string{configDir},
		EnvPrefix:   "",
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
