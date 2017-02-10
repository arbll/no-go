package config

import (
	"github.com/spf13/viper"
	"fmt"
)

func Load(file string)  {
	initConfig(file)
}

func initConfig(file string)  {
	initDefaults()
	viper.SetConfigType("yaml")

	viper.SetConfigFile(file)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("No configuration foud, unsing defaults values")
	}
}

func initDefaults() {
	viper.SetDefault("webserver.port", "8080")
	viper.SetDefault("webserver.directory", "./www/")
	viper.SetDefault("webserver.redirect", "false")
	viper.SetDefault("webserver.redirectUrl", "")
}