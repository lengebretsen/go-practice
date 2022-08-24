package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func setDefaultConfig() {
	//Database
	viper.SetDefault("database.name", "go-practice")
	viper.SetDefault("database.user", "gousr")
	viper.SetDefault("database.pass", "gopass")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")

	//Gin server
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "8080")
}

func LoadConfig() {
	setDefaultConfig()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			println("No config file found, loading with default configs")
		} else {
			panic(fmt.Errorf("fatal error reading config file: %w", err))
		}
	}
}
