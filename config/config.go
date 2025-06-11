package config

import (
	"embed"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

//go:embed config.yaml
var _ embed.FS

var cfg *Config //nolint:gochecknoglobals

func LoadConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "viper read-in-config")
	}

	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)

		if strings.HasPrefix(v, "${") {
			viper.Set(k, os.ExpandEnv(v))
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return errors.Wrap(err, "viper unmarshal config")
	}

	return nil
}

type Config struct {
	Env    string       `mapstructure:"env"`
	Server ServerConfig `mapstructure:"server"`
}

type ServerConfig struct {
	apiPort int `mapstructure:"apiPort"`
}

func GetServerConfig() ServerConfig {
	return cfg.Server
}

func GetEnv() string {
	return cfg.Env
}
