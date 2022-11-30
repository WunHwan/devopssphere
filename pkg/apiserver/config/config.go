package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io.github/devopssphere/pkg/simple/client/db"
	"strings"
)

const (
	defaultConstant string = "config"

	// DefaultConfigurationName is the default name of configuration
	defaultConfigurationFile = defaultConstant

	// defaultConfigurationType is the default content type of configuration
	defaultConfigurationType string = "toml"

	// defaultConfigurationPath is the default location of the configuration services
	defaultConfigurationPath string = "/etc/devopssphere"
)

type Config struct {
	DatabaseOptions *db.Options `mapstructure:"database,omitempty"`
}

func New() *Config {
	return &Config{
		DatabaseOptions: db.NewDatabaseOptions(),
	}
}

func TryLoadFromDisk() (*Config, error) {
	viper.New()

	viper.SetConfigName(defaultConfigurationFile)
	viper.SetConfigType(defaultConfigurationType)
	viper.AddConfigPath(defaultConfigurationPath)

	// Load from current working directory, only used for debuging
	viper.AddConfigPath(".")

	// Load from Environment variables
	viper.SetEnvPrefix(defaultConstant)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "-"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, err
		} else {
			return nil, fmt.Errorf("fatal error config file: %s", err)
		}
	}

	config := New()
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, nil
}
