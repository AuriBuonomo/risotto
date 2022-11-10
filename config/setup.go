package config

import (
	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
)

var configPath = configFolder + "/" + configName
var configFolder = configdir.LocalConfig("risotto")
var configName = "config.toml"

func Setup() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(configFolder)
	err := viper.ReadInConfig()
	if err != nil {
		return createConfig()
	}
	return nil
}
