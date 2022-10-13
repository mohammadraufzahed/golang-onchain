package config

import (
	"os"

	"github.com/spf13/viper"
)

var Config *viper.Viper = viper.New()

func Initialize() {
	configFile := "development.env"
	if os.Getenv("ENV") == "PRODUCTION" {
		configFile = "production.env"
	}
	// Config the viper
	Config.SetConfigFile(configFile)
	Config.AddConfigPath(".")
	// Read the config
	err := Config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
