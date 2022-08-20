package config

import (
	"github.com/spf13/viper"
	"strings"
)

func New[M any](conf *M, path string, configName string, prefix string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(conf); err != nil {
		return err
	}
	return nil
}
