package config

import (
	"os"

	"github.com/spf13/viper"
)

var Viper *viper.Viper = viper.New()

func Initialize() {
	configFile := "development.env"
	if os.Getenv("ENV") == "PRODUCTION" {
		configFile = "production.env"
	}
	// Config the viper
	Viper.SetConfigFile(configFile)
	Viper.AddConfigPath(".")
	// Read the config
	err := Viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
